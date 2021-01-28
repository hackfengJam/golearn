package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// (node Node) 接收者，相当于 this/self
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// (node *Node) 值传递，如果不使用*Node，setValue将失败
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	// 返回一个局部变量的地址
	return &Node{Value: value}

}
