package main

import "fmt"

type AVLTree struct {
	root *Node
}

func newAVLTree(value int) *AVLTree {
	return &AVLTree{&Node{value: value}}
}

// Insert insert a node in a tree
func (t *AVLTree) Insert(value int) {
	t.root = t.root.insert(value)
}

// Delete delete the node with value
func (t *AVLTree) Delete(value int) {
	t.root = t.root.delete(value)
}

// Inorder inorder traversal
func (t *AVLTree) Inorder() {
	t.root.inorder()
}

// BFT breadth first traversal
func (t *AVLTree) BFT() {
	t.root.bft()
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) *Node {
	if n.value > value {
		if n.left != nil {
			n.left.insert(value)
		} else {
			n.left = &Node{value: value}
		}
	} else if n.value < value {
		if n.right != nil {
			n.right.insert(value)
		} else {
			n.right = &Node{value: value}
		}
	}
	return balance(n)
}

func balance(node *Node) *Node {
	if node == nil {
		return node
	}

	bf := getBalanceFactor(node)
	if bf < -1 {
		// right left rotation
		if getBalanceFactor(node.right) > 0 {
			node.right = rightRotation(node.right)
		}
		// left rotation
		return leftRotation(node)
	}

	if bf > 1 {
		// left right rotation
		if getBalanceFactor(node.left) < 0 {
			node.left = leftRotation(node.left)
		}
		// right rotation
		return rightRotation(node)
	}
	return node
}

func getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return getHeight(node.left) - getHeight(node.right)
}

func getHeight(node *Node) int {
	if node == nil {
		return -1
	}
	return 1 + max(getHeight(node.left), getHeight(node.right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func leftRotation(node *Node) *Node {
	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node
	return newRoot
}

func rightRotation(node *Node) *Node {
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node
	return newRoot
}

func (n *Node) delete(value int) *Node {
	switch {
	case n.value > value:
		n.left = n.left.delete(value)
	case n.value < value:
		n.right = n.right.delete(value)
	default:
		// two child
		if n.left != nil && n.right != nil {
			min := fetchMin(n.right)
			n.value = min.value
			n.right = n.right.delete(min.value)
		} else {
			// no child or only one child
			var temp *Node
			if n.left != nil {
				temp = n.left
			} else {
				temp = n.right
			}

			if temp == nil {
				return nil
			} else {
				n = temp
			}
		}
	}
	return balance(n)
}

func fetchMin(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (n *Node) inorder() {
	if n != nil {
		n.left.inorder()
		fmt.Printf("-> %d ", n.value)
		n.right.inorder()
	}
}

func (n *Node) bft() {
	queue := append([]*Node{}, n)
	for len(queue) > 0 {
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
	fmt.Println("insert")
	fmt.Println("left rotation")
	tree := newAVLTree(6)
	tree.Insert(7)
	tree.Insert(8)
	/*
		6                                7
		 \       left rotation (6)      / \
		  7            -->             6   8
		   \
		    8
	*/
	tree.Inorder()
	fmt.Println()
	// -> 6 -> 7 -> 8

	tree.BFT()
	fmt.Println()
	// -> 7 -> 6 -> 8

	fmt.Println("right rotation")
	tree = newAVLTree(7)
	tree.Insert(5)
	tree.Insert(2)
	/*
	       7                          5
	      /    right rotation (7)    / \
	     5            -->           2   7
	    /
	   2
	*/
	tree.Inorder()
	fmt.Println()
	// -> 2 -> 5 -> 7

	tree.BFT()
	fmt.Println()
	// -> 5 -> 2 -> 7

	fmt.Println("right left rotation")
	tree = newAVLTree(6)
	tree.Insert(10)
	tree.Insert(7)
	/*
	   6                            6                              7
	    \     right rotation (10)    \      left rotation (6)     / \
	    10           -->              7            -->           6  10
	    /                              \
	   7                               10
	*/
	tree.Inorder()
	fmt.Println()
	// -> 6 -> 7 -> 10

	tree.BFT()
	fmt.Println()
	// -> 7 -> 6 -> 10

	fmt.Println("left right rotation")
	tree = newAVLTree(10)
	tree.Insert(6)
	tree.Insert(8)
	/*
	    10                           10                               8
	    /      left rotation (6)     /      right rotation (10)      / \
	   6             -->            8              -->              6  10
	    \                          /
	     8                        6
	*/

	tree.Inorder()
	fmt.Println()
	// -> 6 -> 8 -> 10

	tree.BFT()
	fmt.Println()
	// -> 8 -> 6 -> 10

	fmt.Printf("\ndelete\n")
	fmt.Println("left rotation")
	tree = newAVLTree(6)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(8)
	tree.Delete(3)
	/*
		  6                  6
		 / \     delete 3     \
		3   7       -->        7
		     \                  \
		      8                  8

		6                               7
		 \      left rotation (6)      / \
		  7           -->             6   8
		   \
		    8
	*/

	tree.Inorder()
	fmt.Println()
	// -> 6 -> 7 -> 8

	tree.BFT()
	fmt.Println()
	// -> 7 -> 6 -> 8

	fmt.Println("right rotation")
	tree = newAVLTree(7)
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(2)
	tree.Delete(10)
	/*
	       7                       7
	      / \      delete 10      /
	     5   10       -->        5
	    /                       /
	   2                       2

	       7                          5
	      /    right rotation (7)    / \
	     5            -->           2   7
	    /
	   2
	*/

	tree.Inorder()
	fmt.Println()
	// -> 2 -> 5 -> 7

	tree.BFT()
	fmt.Println()
	// -> 5 -> 2 -> 7

	fmt.Println("right left rotation")
	tree = newAVLTree(6)
	tree.Insert(10)
	tree.Insert(2)
	tree.Insert(7)
	tree.Delete(2)
	/*
	   6                        6
	  / \        delete 2        \
	 2  10          -->          10
	    /                        /
	   7                        7

	   6                            6                              7
	    \     right rotation (10)    \      left rotation (6)     / \
	    10           -->              7            -->           6  10
	    /                              \
	   7                               10
	*/

	tree.Inorder()
	fmt.Println()
	// -> 6 -> 10 -> 7

	tree.BFT()
	fmt.Println()
	// -> 7 -> 6 -> 10

	fmt.Println("left right rotation")
	tree = newAVLTree(10)
	tree.Insert(6)
	tree.Insert(15)
	tree.Insert(8)
	tree.Delete(15)
	/*
	    10                      10
	    / \      delete 15      /
	   6  15        -->        6
	    \                       \
	     8                       8

	    10                           10                               8
	    /      left rotation (6)     /      right rotation (10)      / \
	   6             -->            8              -->              6  10
	    \                          /
	     8                        6
	*/

	tree.Inorder()
	fmt.Println()
	// -> 6 -> 8 -> 10

	tree.BFT()
	fmt.Println()
	// -> 8 -> 6 -> 10
}
