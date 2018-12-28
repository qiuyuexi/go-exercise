package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	Size int
	head *Node
	tail *Node
}

func (list *List) Init() {
	(*list).Size = 0
	(*list).head = nil
	(*list).tail = nil
}

func (list *List) append(node *Node) (ret bool, err error) {
	ret = true
	err = nil
	if node == nil {
		err = errors.New("node 未初始化")
		ret = false
		return ret, err
	}
	(*node).next = nil
	if (*list).Size == 0 {
		(*list).Size = 1
		(*list).tail = node
		(*list).head = node
	} else {

		oldTail := (*list).tail
		(*oldTail).next = node
		(*list).tail = node
		(*list).Size++
	}
	return ret, nil
}

func (list *List) remove(index int) (ret bool, err error) {
	ret = true
	err = nil
	if index >= (*list).Size {
		err = errors.New("不存在该项")
		ret = false
		return ret, err
	}

	if index == 0 {
		head := (*list).head
		(*list).head = (*head).next

		//只有一个数据
		if ((*list).Size == 1) {
			(*list).head = nil
			(*list).tail = nil
		}
	} else {
		// 1 2 3 4
		preNode, _ := (*list).getOne(index - 1)
		curNode := (*preNode).next
		(*preNode).next = (*curNode).next

		if index == ((*list).Size - 1) {
			(*list).tail = preNode
		}
	}
	(*list).Size--
	fmt.Println((*list).Size)
	return ret, nil
}

func (list List) output() {
	size := list.Size
	head := list.head
	fmt.Println(size)

	for i := 0; i < size; i++ {
		fmt.Println(i, (*head).data)
		head = (*head).next
	}
}

func (list List) getOne(index int) (node *Node, err error) {
	if index >= list.Size || list.Size <= 0 {
		err = errors.New("不存在该项")
		return nil, err
	}
	head := list.head
	for i := 0; i < index; i++ {
		head = (*head).next
	}
	node = head
	return node, nil
}

/**
 1 -> 2 -> 3 ==》 1<-2<-3    1->nil  保存2   2-> 1  保存3  3->2 结束
 */
func (list *List) reverse() {
	var nextNode, curNode, preNode *Node
	curNode = list.head
	preNode = nil
	for i := 0; i < list.Size; i++ {
		nextNode = curNode.next
		curNode.next = preNode
		preNode = curNode
		curNode = nextNode
	}
	list.head, list.tail = list.tail, list.head
}

func main() {

	//链表
	list := new(List)
	list.Init()
	node := Node{1, nil}
	list.append(&node)
	//list.output()

	node1 := Node{3, nil}
	list.append(&node1)
	//list.output()

	node2 := Node{4, nil}

	list.append(&node2)
	list.output()

	//lastOne, _ := list.getOne(1)
	//fmt.Println((*lastOne).data)

	list.reverse()
	list.output()
	list.reverse()
	list.output()
}
