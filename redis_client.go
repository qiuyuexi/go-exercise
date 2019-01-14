package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const CRLF string = "\r\n"

var pool *redisClientPool
var poolInit sync.Once

/**
redis 连接池
 */
type redisClientPool struct {
	sync.RWMutex
	pool     map[string]*redisClient
	usedPool map[string]map[string]*redisClient
	freePool map[string]map[string]*redisClient
}

/*
单例
*/
func getPool() *redisClientPool {
	poolInit.Do(func() {
		pool = &redisClientPool{pool: map[string]*redisClient{}}
	})
	return pool
}

/**
获取连接
 */
func (redisClientPool *redisClientPool) getClient(network string, address string) *redisClient {
	redisClientPool.RWMutex.Lock()
	defer redisClientPool.RWMutex.Unlock()
	key := address + network
	if redisClientPool.pool[key] == nil || redisClientPool.pool[key].expireTimeStamp < time.Now().Unix() {
		redis := &redisClient{}
		err := redis.init(network, address)
		if err != nil {
			redisClientPool.pool[key] = nil
		} else {
			redisClientPool.pool[key] = redis
		}

		fmt.Println("connect init")
	}

	return redisClientPool.pool[key]

}

/**
关闭所有连接
 */
func (redisClientPool *redisClientPool) close() {
	for _, v := range redisClientPool.pool {
		v.conn.Close()
	}
}

/**
redis 客户端连接
 */
type redisClient struct {
	conn             net.Conn
	network, address string
	mutex            sync.Mutex
	bufr             *bufio.Reader
	bufw             *bufio.Writer
	writeTimeOut     time.Duration
	readTimeOut      time.Duration
	expireTimeStamp  int64
}

type v struct {
	len int
	v   []byte
}

func (redis *redisClient) send(cmd string) {
	cmd = getCmd(cmd)
	_, err := redis.bufw.WriteString(cmd)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = redis.bufw.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (redis *redisClient) receive() interface{} {
	b := make([]byte, 1024)
	_, err := redis.bufr.Read(b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(b))
	var result interface{}

	switch b[0] {
	case '+', '-':
		result = string(b[1:])
		fmt.Println("result:" + string(b[1:]))
	case ':':
		result = string(b[1 : len(b)-2])
	case '$':
		var index int
		b = b[1:]
		index = getLineLen(b)
		length, _ := strconv.Atoi((string(b[0:index])))
		b = b[index+2:]
		index = getLineLen(b)
		line := b[0:index]
		value := &v{length, line}
		result = value
	case '*':
		var index int
		b = b[1:]
		index = getLineLen(b)
		num, _ := strconv.Atoi(string(b[0:index]))
		b = b[index:]
		value := make([]*v, num)
		for i := 0; i < num; i++ {
			b = b[1:]
			index = getLineLen(b)
			length, _ := strconv.Atoi(string(b[0:index]))

			b = b[index+2:]
			index = getLineLen(b)
			str := b[0:index]
			value[i] = &v{length, str}
			b = b[index+2:]
		}
		result = value
	}
	return result
}

/**
初始化
 */
func (redis *redisClient) init(network string, address string) error {
	var err error
	redis.address = address
	redis.network = network
	conn, err := net.DialTimeout(redis.network, redis.address, 1*time.Minute)
	if err != nil {
		fmt.Println("tcp connect error:" + err.Error())
		return err
	}
	redis.conn = conn
	redis.bufr = bufio.NewReader(conn)
	redis.bufw = bufio.NewWriter(conn)
	redis.expireTimeStamp = time.Now().Unix() + 60
	return nil
}

/**
	按照reids 协议，拼接命令
	redis 通信协议
	http://redisdoc.com/topic/protocol.html?spm=a2c4e.11153940.blogcont651002.9.551031dan9fspp
 */
func getCmd(input string) string {
	input = strings.Replace(input, "\n", "", -1)

	inputByte := bytes.NewBufferString("*")
	inputList := strings.Split(input, " ")
	inputByte.WriteString(strconv.Itoa(len(inputList)))
	inputByte.WriteString(CRLF)

	for _, v := range inputList {
		inputByte.WriteString("$")
		inputByte.WriteString(strconv.Itoa(len(v)))
		inputByte.WriteString(CRLF)
		inputByte.WriteString(v)
		inputByte.WriteString(CRLF)
	}
	//fmt.Println(inputByte.String())
	return inputByte.String()
}

/**
 获取一行的长度， 每一行根据\r\n结尾
 */
func getLineLen(string []byte) int {
	stringLen := len(string)
	index := 0
	for i := 0; i < stringLen; i++ {
		if string[i] == '\r' && (i+1 < stringLen && string[i+1] == '\n') {
			break
		}
		index++
	}
	return index
}

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("main error")
			fmt.Println(err)
		}
	}()
	for {
		input, err := inputReader.ReadString('\n')
		if err == nil {
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println("error")
						fmt.Println(err)
					}
				}()
				redis := getPool().getClient("tcp", "127.0.0.1:6379")
				redis.send(input)
				redis.receive()
			}()
		}
	}
	defer getPool().close()
}
