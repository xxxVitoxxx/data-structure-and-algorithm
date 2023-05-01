package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value       int
	left, right *Node
}

type BinaryTree struct {
	root *Node
}

func createNode(value int) *Node {
	return &Node{value: value}
}

func newBinaryTree(value int) *BinaryTree {
	return &BinaryTree{
		root: &Node{value: value},
	}
}

// GetNodeCount count the number of nodes in tree
func (root *Node) GetNodeCount() int {
	if root == nil {
		return 0
	}
	return root.left.GetNodeCount() + root.right.GetNodeCount() + 1
}

// Preorder to traversal a tree in preorder
// visit the root
// traversal the left subtree of the root
// traversal the right subtree of the root
func (root *Node) Preorder() {
	if root != nil {
		fmt.Printf("-> %v ", root.value)
		root.left.Preorder()
		root.right.Preorder()
	}
}

// Postorder to traversal a tree in postorder traversal
// traversal the left subtree of the root
// traversal the right subtree of the root
// visit the root
func (root *Node) Postorder() {
	if root != nil {
		root.left.Postorder()
		root.right.Postorder()
		fmt.Printf("-> %v ", root.value)
	}
}

// Inorder to traversal a tree in inorder traversal
// traversal the left subtree of the root
// visit the root
// traversal the right subtree of the root
func (root *Node) Inorder() {
	if root != nil {
		root.left.Inorder()
		fmt.Printf("-> %v ", root.value)
		root.right.Inorder()
	}
}

// LayerOrder output the value of node following left to right and layer by layer
func (root *Node) LayerOrder() {
	if root == nil {
		return
	}

	var queue []*Node
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("-> %v ", node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

// GetTreeDegree count degree of a tree
func (root *Node) GetTreeDegree() int {
	if root != nil {
		var maxDegree int
		queue := append([]*Node{}, root)
		for len(queue) != 0 {
			element := queue[0]
			queue = queue[1:]

			var degree int
			if element.left != nil {
				degree += 1
				queue = append(queue, element.left)
			}

			if element.right != nil {
				degree += 1
				queue = append(queue, element.right)
			}

			if degree > maxDegree {
				maxDegree = degree
			}
		}
		return maxDegree
	}
	return 0
}

// Search search the node that stores the value
func (root *Node) Search(value int) (*Node, error) {
	if root != nil {
		queue := append([]*Node{}, root)
		for len(queue) != 0 {
			element := queue[0]
			queue = queue[1:]
			if element.value == value {
				return element, nil
			}

			if element.left != nil {
				queue = append(queue, element.left)
			}

			if element.right != nil {
				queue = append(queue, element.right)
			}
		}
	}
	return &Node{}, errors.New("the value not exist in node of a tree")
}

func main() {
	tree := newBinaryTree(5)
	tree.root.left = createNode(1)
	tree.root.right = createNode(8)
	tree.root.left.left = createNode(6)
	tree.root.left.right = createNode(9)
	tree.root.right.left = createNode(3)
	tree.root.right.right = createNode(10)
	tree.root.right.right.left = createNode(7)
	/*
	        5
	      /   \
	     1     8
	    / \   / \
	   6   9 3   10
	            /
	           7
	*/

	degree := tree.root.GetTreeDegree()
	fmt.Println("degree of a tree", degree)

	count := tree.root.GetNodeCount()
	fmt.Println("count of nodes in tree: ", count)
	// 8

	node, err := tree.root.Search(7)
	if err != nil {
		fmt.Println("search err: ", err)
	}
	fmt.Printf("node %v exist in tree\n", node)
	// node &{7 <nil> <nil>} exist in tree

	tree.root.Preorder()
	fmt.Println()
	// -> 5 -> 1 -> 6 -> 9 -> 8 -> 3 -> 10 -> 7

	tree.root.Postorder()
	fmt.Println()
	// -> 6 -> 9 -> 1 -> 3 -> 7 -> 10 -> 8 -> 5

	tree.root.Inorder()
	fmt.Println()
	// -> 6 -> 1 -> 9 -> 5 -> 3 -> 8 -> 7 -> 10

	tree.root.LayerOrder()
	fmt.Println()
	// -> 5 -> 1 -> 8 -> 6 -> 9 -> 3 -> 10 -> 7
}
