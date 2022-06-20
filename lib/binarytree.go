package lib

import (
	"encoding/json"
	"fmt"
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
	}else {
		fmt.Println("key already exists")
	}
}
// Search
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if n.Key < key {
		// move right
		return n.Right.Search(key)
	}else if n.Key > key {
		// move left
		return n.Left.Search(key)
	}
	// it is found
	return true
}

func BinaryTreeMain() {
	tree := &Node{Key: 100}
	for index:=0; index<200; index++ {
		tree.Insert(index)
	}
	data, err := json.Marshal(tree)
	HandleError(err)
	err = ioutil.WriteFile("./DB/tree.json", data, 0755)
	HandleError(err)
	fmt.Println(tree.Search(141))
	fmt.Println(tree.Search(241))
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}