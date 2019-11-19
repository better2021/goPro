package study

import "fmt"

type treeNode struct {
	value int
	left , right *treeNode
}

func Four()  {
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(10)

	nodes := []treeNode{
		{value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)
}

func createNode(value int) *treeNode{
	return &treeNode{value:value}
}