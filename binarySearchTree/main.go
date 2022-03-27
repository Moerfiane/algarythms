package main

import "fmt"

var count int

//define node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//insert nodes into the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}

}

//search node will take in  akey value and return true if there is a node with that value
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(27)
	fmt.Println(tree)

	fmt.Println(tree.Search(76))
	fmt.Println(count)
	fmt.Println(tree.Search(400))
	fmt.Println(count)
	fmt.Println(tree.Search(310))
	fmt.Println(count)
}
