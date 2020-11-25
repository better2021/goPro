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

	var r Retiever
	fmt.Println(download(r))
}

func createNode(value int) *treeNode{
	return &treeNode{value:value}
}

/*
接口类型 interface
*/

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string{
	return r.Get("www.imooc.com")
}

