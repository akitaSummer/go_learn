package main

import (
	"fmt"
	"learn_go/struct/tree"
)

// 扩展包
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func (myNode *myTreeNode) postOrderFunc(f func(node *myTreeNode)) {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{node: myNode.node.Left}
	left.postOrderFunc(f)
	right := myTreeNode{node: myNode.node.Right}
	right.postOrderFunc(f)
	f(myNode)
}

func main() {
	var root tree.Node
	fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	node := tree.CreateNode(2)

	fmt.Println(node)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Print()
	root.SetValue(100)
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	pRoot.Traverse()

	myNode := myTreeNode{&root}
	myNode.postOrder()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("max node value:", maxNode)
}
