package main

import (
	"errors"
	"fmt"
)

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

// GetNodeCount count the number of nodes in tree
func (node *Node) GetNodeCount() int {
	if node == nil {
		return 0
	}
	return node.left.GetNodeCount() + node.right.GetNodeCount() + 1
}

// GetTreeDegree count degree of a tree
func (node *Node) GetTreeDegree() int {
	if node != nil {
		var maxDegree int
		queue := append([]*Node{}, node)
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

// Search return the node that stores the value
func (node *Node) Search(value int) (*Node, error) {
	if node != nil {
		queue := append([]*Node{}, node)
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

// DeleteNode delete a node
func (node *Node) DeleteNode(value int) {
	node = delete(node, value)
}

func delete(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	switch {
	case node.value > value:
		node.left = delete(node.left, value)
	case node.value < value:
		node.right = delete(node.right, value)
	default:
		if node.left == nil && node.right == nil {
			return nil
		}

		if node.left == nil {
			return node.right
		}

		if node.right == nil {
			return node.left
		}

		// two child
		max := fetchMax(node.left)
		node.value = max.value
		node.left = delete(node.left, max.value)
	}

	return node
}

func fetchMax(node *Node) *Node {
	for node.right != nil {
		node = node.right
	}
	return node
}

// FetchMin _
func (node *Node) FetchMin() int {
	if node.left == nil {
		return node.value
	}
	return node.left.FetchMin()
}

// FetchMax _
func (node *Node) FetchMax() int {
	if node.right == nil {
		return node.value
	}
	return node.right.FetchMax()
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

	count := tree.root.GetNodeCount()
	fmt.Println("count of nodes in a tree: ", count)
	// count of nodes in a tree:  8

	degree := tree.root.GetTreeDegree()
	fmt.Println("degree of a tree: ", degree)
	// degree of a tree:  2

	min := tree.root.FetchMin()
	fmt.Println("min: ", min)
	// min:  1

	max := tree.root.FetchMax()
	fmt.Println("max: ", max)
	// max:  12

	node, err := tree.root.Search(12)
	if err != nil {
		fmt.Println("search err: ", err)
	}
	fmt.Printf("node %v exist in a tree\n", node)
	// node &{12 0x1400011a0a8 <nil>} exist in a tree

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

	// delete a node that has a child
	fmt.Println("delete 12")
	tree.root.DeleteNode(12)
	/*
		       5
		     /   \
		   3      8
		 /   \   /  \
		1     4 7    9
	*/

	fmt.Println("inorder")
	tree.root.Inorder()
	fmt.Println()
	// -> 1 -> 3 -> 4 -> 5 -> 7 -> 8 -> 9

	fmt.Println("BFT")
	tree.root.BFT()
	fmt.Println()
	// -> 5 -> 3 -> 8 -> 1 -> 4 -> 7 -> 9

	// delete a node that is a leaf
	fmt.Println("delete 4")
	tree.root.DeleteNode(4)
	/*
		       5
		     /   \
		   3      8
		 /       /  \
		1       7    9
	*/

	fmt.Println("inorder")
	tree.root.Inorder()
	fmt.Println()
	// -> 1 -> 3 -> 5 -> 7 -> 8 -> 9

	fmt.Println("BFT")
	tree.root.BFT()
	fmt.Println()
	// -> 5 -> 3 -> 8 -> 1 -> 7 -> 9

	// delete a node that has two child
	fmt.Println("delete 8")
	tree.root.DeleteNode(8)
	/*
		       5
		     /   \
		   3      7
		 /          \
		1            9
	*/

	fmt.Println("inorder")
	tree.root.Inorder()
	fmt.Println()
	// -> 1 -> 3 -> 5 -> 7 -> 9

	fmt.Println("BFT")
	tree.root.BFT()
	fmt.Println()
	// -> 5 -> 3 -> 7 -> 1 -> 9
}
