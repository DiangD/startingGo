package tree

import "fmt"

//为结构定义的方法要在一个包内，可以是不同的文件

func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

//使用函数作为参数，更灵活拓展方法的功能
func (node *Node) TraverseFunc(f func(node *Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	c := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			c <- node
		})
		close(c)
	}()
	return c
}
