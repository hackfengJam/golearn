package main

import (
	"../../tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	// 包装一下
	//myTreeNode{node: myNode.node.Left}.postOrder()
	//myTreeNode{myNode.node.Right}.postOrder()
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	//root.print()
	//root.setValue(100)
	//root.print()
	//
	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()
	//
	//var pRoot2 *treeNode
	//
	//// nil 可以调用方法
	//pRoot2.setValue(200)
	//pRoot2 = &root
	//pRoot2.setValue(300)
	//pRoot2.print()

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
