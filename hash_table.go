package main

import (
	"fmt"
)

/**
数组+链表
map 本身就是hashtable
 */

type hashTableNode struct {
	key  interface{}
	data interface{}
	next *hashTableNode
	pre  *hashTableNode
}

type hashTableList struct {
	size int
	head *hashTableNode
	tail *hashTableNode
}

/**
新增节点
 */
func (hashTableList *hashTableList) append(key interface{}, data interface{}) {
	node := &hashTableNode{key, data, nil, nil}
	if hashTableList.size == 0 {
		hashTableList.head = node
		hashTableList.tail = node
		hashTableList.size = 0
	} else {
		oldNode := hashTableList.tail
		oldNode.next = node
		node.pre = oldNode
		hashTableList.tail = node
	}
	hashTableList.size++
}

type HashTable struct {
	key [5]*hashTableList
}


func (hashTable *HashTable) hash(key string) int {
	return 0
}

func (hashTable *HashTable) put(key string, value interface{}) {
	index := hashTable.hash(key)
	if hashTable.key[index] == nil {
		hashTable.key[index] = &hashTableList{0, nil, nil}
	}
	hashTable.key[index].append(key, value)
}

func (hashTable *HashTable) remove(key string) bool {
	index := hashTable.hash(key)
	list := hashTable.key[index]
	node := list.head
	for i := 0; i < list.size; i++ {
		if node.key == key {
			//移除该节点
			node.pre.next = node.next;
			node.next.pre = node.pre;
			list.size--
			return true
		}
		node = node.next
	}
	return false
}

func (hashTable *HashTable) get(key string) interface{} {
	index := hashTable.hash(key)
	list := hashTable.key[index]
	node := list.head
	for i := 0; i < list.size; i++ {
		if node.key == key {
			return node.data
		}
		node = node.next
	}
	return nil
}

func main() {
	table := new(HashTable)
	//table.init()
	table.put("1", 2)
	table.put("2", []int{1, 2, 3})
	table.put("3", 4)
	table.put("4", map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"})
	fmt.Println(table.get("3"))
	table.remove("3")
	fmt.Println(table.get("3"))
	data := table.get("2")

	arr, ok := data.([]int)
	if ok {
		fmt.Println(arr)
	}

	mm := table.get("4");
	m, ok := mm.(map[string]string)
	if ok {
		fmt.Println(m)
	}

}
