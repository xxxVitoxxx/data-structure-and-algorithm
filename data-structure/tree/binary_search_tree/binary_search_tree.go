package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

type BST struct {
	root *Node
}

func newBST(value int) *BST {
	return &BST{root: &Node{value: value}}
}

// Insert insert an element in a tree
func (node *Node) Insert(element int) {
	if node != nil {
		if node.value > element {
			if node.left == nil {
				node.left = &Node{value: element}
			} else {
				node.left.Insert(element)
			}
		} else {
			if node.right == nil {
				node.right = &Node{value: element}
			} else {
				node.right.Insert(element)
			}
		}
	}
}

// Preorder traversal a tree in preorder
func (node *Node) Preorder() {
	if node != nil {
		fmt.Printf("-> %v ", node.value)
		node.left.Preorder()
		node.right.Preorder()
	}
}

// Inorder traversal a tree in inorder
func (node *Node) Inorder() {
	if node != nil {
		node.left.Inorder()
		fmt.Printf("-> %v ", node.value)
		node.right.Inorder()
	}
}

// Postorder traversal a tree in postorder
func (node *Node) Postorder() {
	if node != nil {
		node.left.Postorder()
		node.right.Postorder()
		fmt.Printf("-> %v ", node.value)
	}
}

// BFT breadth first traversal
func (node *Node) BFT() {
	queue := append([]*Node{}, node)
	for len(queue) != 0 {
		element := queue[0]
		queue = queue[1:]
		fmt.Printf("-> %d ", element.value)

		if element.left != nil {
			queue = append(queue, element.left)
		}

		if element.right != nil {
			queue = append(queue, element.right)
		}
	}
}

func main() {
	tree := newBST(5)
	tree.root.Insert(8)
	tree.root.Insert(3)
	tree.root.Insert(1)
	tree.root.Insert(4)
	tree.root.Insert(12)
	tree.root.Insert(9)
	tree.root.Insert(7)

	/*
		       5
		     /   \
		   3      8
		 /   \   /  \
		1     4 7   12
		            /
		           9
	*/

	tree.root.Preorder()
	fmt.Println()
	// -> 5 -> 3 -> 1 -> 4 -> 8 -> 7 -> 12 -> 9

	tree.root.Inorder()
	fmt.Println()
	// -> 1 -> 3 -> 4 -> 5 -> 7 -> 8 -> 9 -> 12

	tree.root.Postorder()
	fmt.Println()
	// -> 1 -> 4 -> 3 -> 7 -> 9 -> 12 -> 8 -> 5

	tree.root.BFT()
	fmt.Println()
	// -> 5 -> 3 -> 8 -> 1 -> 4 -> 7 -> 12 -> 9
}
