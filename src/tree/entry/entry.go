package main

import (
	"fmt"
	"tree"
)

func main() {
	root := tree.Node{
		Value: 3,
	}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	//内建函数，返回地址
	//不论地址还是结构一律使用“.”来访问成员，同于c的“->”
	root.Right.Left = new(tree.Node)

	nodes := []tree.Node{
		{},
		{6, nil, tree.NewTreeNode(1, nil, nil)},
		{7, nil, &root},
	}
	fmt.Println(nodes)
	root.Print()
	root.SetValue(100)
	root.Print()

	var pRoot *tree.Node
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()
	root.Traverse()
}
