package main

import (
	"fmt"
)

type Node struct {
	value       int
	left, right *Node
}

type BST struct {
	root *Node
}

func NewBST(value int) *BST {
	return &BST{
		root: &Node{value: value},
	}
}

func (node *Node) Insert(value int) {
	if node != nil {
		if node.value > value {
			if node.left != nil {
				node.left.Insert(value)
			} else {
				node.left = &Node{value: value}
			}
		}

		if node.value < value {
			if node.right != nil {
				node.right.Insert(value)
			} else {
				node.right = &Node{value: value}
			}
		}
	}
}

func isBalance(node *Node) bool {
	if node == nil {
		return true
	}

	queue := append([]*Node{}, node)
	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		leftHeight := height(element.left)
		rightHeight := height(element.right)
		if abs(leftHeight, rightHeight) > 1 {
			return false
		}

		if element.left != nil {
			queue = append(queue, element.left)
		}

		if element.right != nil {
			queue = append(queue, element.right)
		}
	}
	return true
}

func height(node *Node) int {
	if node == nil {
		return -1
	}
	return 1 + max(height(node.left), height(node.right))
}

func abs(nu1, nu2 int) int {
	if nu1-nu2 > 0 {
		return nu1 - nu2
	}
	return nu2 - nu1
}

func max(nu1, nu2 int) int {
	if nu1 > nu2 {
		return nu1
	}
	return nu2
}

func main() {
	tree := NewBST(8)
	tree.root.Insert(5)
	tree.root.Insert(6)
	tree.root.Insert(13)
	tree.root.Insert(10)
	tree.root.Insert(12)
	tree.root.Insert(17)
	/*
	      8
	    /   \
	   5     13
	    \    / \
	     6  10  17
	          \
	          12
	*/

	fmt.Println(isBalance(tree.root))
	// true

	tree = NewBST(9)
	tree.root.Insert(11)
	tree.root.Insert(3)
	tree.root.Insert(1)
	tree.root.Insert(4)
	tree.root.Insert(7)
	/*
	       9
	      / \
	     3   11
	    / \
	   1   4
	        \
	         7
	*/

	fmt.Println(isBalance(tree.root))
	// false
}
