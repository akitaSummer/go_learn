package tree

import "fmt"

// 包中使用大写开头为public，小写开头为private
type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

// go中参数都是值传递，为能修改其属性，需要传递指针
func (node *Node) SetValue(value int) {
	// go中nil指针可以调用方法
	if node == nil {
		fmt.Println("setting value to nil")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//func main() {
//	var root treeNode
//	fmt.Println(root)
//
//	root = treeNode{value: 3}
//	root.Left = &treeNode{}
//	root.Right = &treeNode{5, nil, nil}
//	root.Right.Left = new(treeNode)
//
//	node := createTreeNode(2)
//
//	fmt.Println(node)
//
//	nodes := []treeNode{
//		{value: 3},
//		{},
//		{6, nil, &root},
//	}
//	fmt.Println(nodes)
//
//	root.print()
//	root.setValue(100)
//	root.print()
//
//	pRoot := &root
//	pRoot.print()
//	pRoot.setValue(200)
//	pRoot.print()
//
//	pRoot.Traverse()
//}
