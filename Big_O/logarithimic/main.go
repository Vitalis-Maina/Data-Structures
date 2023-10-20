package main

import "fmt"

type Node struct {
	Left  *Node
	Key   int
	Right *Node
}

// Insert operation
func Insert(n *Node, k int) *Node {
	if n == nil {
		return &Node{Key: k}
	}
	if n.Key < k {
		// Move right
		n.Right = Insert(n.Right, k)
	} else if n.Key > k {
		// Move left
		n.Left = Insert(n.Left, k)
	}
	return n
}

// Search operation
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		// Move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// Move left
		return n.Left.Search(k)
	}
	return true
}

func main() {

	var binaryTree *Node // Initialize root as nil
	fmt.Println(binaryTree)
	binaryTree = Insert(binaryTree, 34)
	binaryTree = Insert(binaryTree, 4)
	binaryTree = Insert(binaryTree, 134)
	binaryTree = Insert(binaryTree, 78)
	binaryTree = Insert(binaryTree, 89)
	binaryTree = Insert(binaryTree, 52)
	binaryTree = Insert(binaryTree, 14)
	binaryTree = Insert(binaryTree, 178)
	fmt.Println(binaryTree)
	fmt.Println(binaryTree.Search(52))

}
