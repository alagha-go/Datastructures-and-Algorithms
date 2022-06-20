package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(400)
	data, err := json.Marshal(tree)
	HandleError(err)
	err = ioutil.WriteFile("./DB/tree.json", data, 0755)
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}