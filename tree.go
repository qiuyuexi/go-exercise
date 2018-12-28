package main

import "fmt"

type Tree struct {
	data        int
	left, right *Tree
}

func (tree *Tree) Init() {

	var data int
	fmt.Scanf("%d", &data)

	if data == 0 {
		tree.data = data
		tree.left = nil
		tree.right = nil
		return;
	} else {
		tree.data = data
		tree.left = new(Tree)
		tree.left.Init()
		tree.right = new(Tree)
		tree.right.Init()
	}
}

func (tree *Tree) Preorder() {
	if tree.data != 0 {
		fmt.Print(tree.data)
		tree.left.Preorder()
		tree.right.Preorder()
	}
	return
}
func (tree *Tree) Midorder() {
	if tree.data != 0 {
		tree.left.Midorder()
		fmt.Print(tree.data)
		tree.right.Midorder()
	}
	return
}

func (tree *Tree) Lastorder() {
	if tree.data != 0 {
		tree.left.Lastorder()
		tree.right.Lastorder()
		fmt.Print(tree.data)
	}
	return
}

func treeInit(tree *Tree) {
	var data int
	fmt.Scanf("%d", &data)
	if data == 0 {
		tree.data = data
		tree.left = nil
		tree.right = nil
	} else {
		tree.data = data
		tree.left = new(Tree)
		tree.right = new(Tree)

		treeInit(tree.left)
		treeInit(tree.right)
	}
	return
}

func treePreorder(tree *Tree) {
	if tree.data != 0 {
		fmt.Print(tree.data)
		treePreorder(tree.left)
		treePreorder(tree.right)

	}
	return
}

func treeMidorder(tree *Tree) {
	if tree.data != 0 {
		treeMidorder(tree.left)
		fmt.Print(tree.data)
		treeMidorder(tree.right)
	}
	return
}

func treeLastorder(tree *Tree) {
	if tree.data != 0 {
		treeLastorder(tree.left)
		treeLastorder(tree.right)
		fmt.Print(tree.data)
	}
	return
}

func main() {
	//二叉树
	//1,2,0,34,5,6,7,8
	tree := new(Tree)
	(*tree).Init()
	(*tree).Preorder()
	fmt.Println("")
	(*tree).Midorder()
	fmt.Println("")
	(*tree).Lastorder()
	fmt.Println("")

	tree1 := new(Tree)
	treeInit(tree1)
	treePreorder(tree1)
	fmt.Println("")
	treeMidorder(tree1)
	fmt.Println("")
	treeLastorder(tree1)
	fmt.Println("")
}
