package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//go语言不存在构造函数~，可以自定义工厂函数
func NewTreeNode(value int, left *Node, right *Node) *Node {
	//返回了局部变量的地址
	return &Node{Value: value, Left: left, Right: right}
}

//值/指针接收者均可以接收值/指针
//区别在于值拷贝还是传指针
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}

//只有指针才可以改变结构内容
//nil指针也可以调用方法！
func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
