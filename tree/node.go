package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//go语言不存在构造函数~，可以自定义工厂函数
func newTreeNode(value int, left *treeNode, right *treeNode) *treeNode {
	//返回了局部变量的地址
	return &treeNode{value: value, left: left, right: right}
}

//值/指针接收者均可以接收值/指针
//区别在于值拷贝还是传指针
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value
}

//只有指针才可以改变结构内容
//nil指针也可以调用方法！
func (node *treeNode) print() {
	fmt.Print(node.value, " ")
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	root := treeNode{
		value: 3,
	}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	//内建函数，返回地址
	//不论地址还是结构一律使用“.”来访问成员，同于c的“->”
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{},
		{6, nil, newTreeNode(1, nil, nil)},
		{7, nil, &root},
	}
	fmt.Println(nodes)
	root.print()
	root.setValue(100)
	root.print()

	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()
	root.traverse()
}
