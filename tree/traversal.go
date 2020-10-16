package tree

//为结构定义的方法要在一个包内，可以是不同的文件

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
