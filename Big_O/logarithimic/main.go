package main

import "fmt"

type Node struct {
	Left  *Node
	Key   int
	Right *Node
}

//insert operation

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

func (n *Node) Search(k int) bool {
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

	binaryTree := &Node{Key: 100}
	fmt.Println(binaryTree)
	binaryTree.Insert(34)
	binaryTree.Insert(4)
	binaryTree.Insert(134)
	binaryTree.Insert(78)
	binaryTree.Insert(89)
	binaryTree.Insert(52)
	binaryTree.Insert(14)
	binaryTree.Insert(178)
	fmt.Println(binaryTree)
	fmt.Println(binaryTree.Search(100))

}
