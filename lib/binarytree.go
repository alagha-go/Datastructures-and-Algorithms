package lib

import "fmt"

// Node
type Node struct {
	Key int							`json:"key,omitempty"`
	Left *Node						`json:"left,omitempty"`
	Right *Node						`json:"right,omitempty"`
}
// insert will add a node to the tree
func (n *Node) Insert(key int) {
	if n.Key < key {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: key}
		}else {
			n.Right.Insert(key)
		}
	}else if n.Key > key{
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: key}
		}else {
			n.Left.Insert(key)
		}
	}
}
// Search


func BinaryTreeMain() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(50)
}