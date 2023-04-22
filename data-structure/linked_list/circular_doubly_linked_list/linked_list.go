package main

import (
	"errors"
	"fmt"
)

type Node struct {
	previous *Node
	data     int
	next     *Node
}

type LinkedList struct {
	head, tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

/*
	method
	1. Insertion − Adds an element at the beginning of the list.
	2. Deletion − Deletes an element at the beginning of the list.
	3. Insert Last − Adds an element at the end of the list.
	4. Delete Last − Deletes an element from the end of the list.
	5. Insert After − Adds an element after an item of the list.
	6. Delete − Deletes an element from the list using the key.
	7. Display forward − Displays the complete list in a forward manner.
	8. Display backward − Displays the complete list in a backward manner.
*/

// Insertion add an element at the beginning of the list
func (ll *LinkedList) Insertion(element int) {
	node := &Node{data: element}
	if ll.head == nil {
		node.previous, node.next = node, node
		ll.head, ll.tail = node, node
		return
	}

	node.previous, node.next = ll.tail, ll.head
	ll.head.previous = node
	ll.tail.next = node
	ll.head = node
}

// InsertLast add an element at the end of the list
func (ll *LinkedList) InsertLast(element int) {
	node := &Node{data: element}
	if ll.head == nil {
		node.previous, node.next = node, node
		ll.head, ll.tail = node, node
		return
	}

	node.previous, node.next = ll.tail, ll.head
	ll.tail.next = node
	ll.head.previous = node
	ll.tail = node
}

// InsertAfter add en element after an item of the list
func (ll *LinkedList) InsertAfter(prev, element int) error {
	if ll.head == nil {
		return errors.New("linked list is empty")
	}

	node := &Node{data: element}
	current := ll.head
	for {
		if current.data == prev {
			if current == ll.tail {
				node.previous, node.next = ll.tail, ll.head
				ll.tail.next = node
				ll.head.previous = node
				ll.tail = node
				return nil
			}

			node.previous, node.next = current, current.next
			current.next.previous = node
			current.next = node
			return nil
		}

		if current == ll.tail {
			break
		}
		current = current.next
	}

	return errors.New("previous data not found")
}

// Deletion delete an element at the beginning of the list
func (ll *LinkedList) Deletion() error {
	if ll.head == nil {
		return errors.New("linked list is empty")
	}

	if ll.head == ll.tail {
		ll.head, ll.tail = nil, nil
		return nil
	}

	ll.head = ll.head.next
	ll.head.previous = ll.tail
	ll.tail.next = ll.head
	return nil
}

// DeleteLast delete an element at from the end of the list
func (ll *LinkedList) DeleteLast() error {
	if ll.tail == nil {
		return errors.New("linked list is empty")
	}

	if ll.head == ll.tail {
		ll.head, ll.tail = nil, nil
		return nil
	}

	ll.tail = ll.tail.previous
	ll.tail.next = ll.head
	ll.head.previous = ll.tail
	return nil
}

// Delete delete an element from the list using the key
func (ll *LinkedList) Delete(key int) error {
	if ll.head == nil {
		return errors.New("linked list is empty")
	}

	current := ll.head
	for {
		if current.data == key {
			if current == ll.head {
				if current == ll.tail {
					ll.head, ll.tail = nil, nil
					return nil
				}

				ll.head = ll.head.next
				ll.head.previous = ll.tail
				ll.tail.next = ll.head
				return nil
			}

			if current == ll.tail {
				ll.tail = ll.tail.previous
				ll.tail.next = ll.head
				ll.head.previous = ll.tail
				return nil
			}

			current.previous.next = current.next
			current.next.previous = current.previous
			return nil
		}

		if current == ll.tail {
			break
		}
		current = current.next
	}
	return errors.New("key not found")
}

// DisplayForward display the complete list in a forward manner
func (ll *LinkedList) DisplayForward() {
	if ll.head == nil {
		return
	}

	current := ll.head
	for {
		fmt.Printf(" -> %v", current.data)
		if current == ll.tail {
			break
		}
		current = current.next
	}
}

// DisplayBackward display the complete list in a backward manner
func (ll *LinkedList) DisplayBackward() {
	if ll.head == nil {
		return
	}

	stack := []int{}
	current := ll.head
	for {
		stack = append(stack, current.data)
		if current == ll.tail {
			break
		}
		current = current.next
	}

	for len(stack) != 0 {
		fmt.Printf(" -> %v", stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
}

func main() {
	ll := NewLinkedList()
	fmt.Println("insertion")
	ll.Insertion(7)
	ll.Insertion(5)
	ll.Insertion(9)

	ll.DisplayForward()
	fmt.Println()
	// -> 9 -> 5 -> 7

	ll.DisplayBackward()
	fmt.Println()
	// -> 7 -> 5 -> 9

	fmt.Println()
	fmt.Println("insert last")
	ll.InsertLast(1)

	ll.DisplayForward()
	fmt.Println()
	// -> 9 -> 5 -> 7 -> 1

	ll.DisplayBackward()
	fmt.Println()
	// -> 1 -> 7 -> 5 -> 9

	fmt.Println()
	fmt.Println("insert after")
	err := ll.InsertAfter(7, 4)
	if err != nil {
		fmt.Println("insert after err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 9 -> 5 -> 7 -> 4 -> 1

	ll.DisplayBackward()
	fmt.Println()
	// -> 1 -> 4 -> 7 -> 5 -> 9

	fmt.Println()
	fmt.Println("deletion")
	err = ll.Deletion()
	if err != nil {
		fmt.Println("deletion err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 5 -> 7 -> 4 -> 1

	ll.DisplayBackward()
	fmt.Println()
	// -> 1 -> 4 -> 7 -> 5

	fmt.Println()
	fmt.Println("delete last")
	err = ll.DeleteLast()
	if err != nil {
		fmt.Println("delete last err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 5 -> 7 -> 4

	ll.DisplayBackward()
	fmt.Println()
	// -> 4 -> 7 -> 5

	fmt.Println()
	fmt.Println("delete")
	err = ll.Delete(4)
	if err != nil {
		fmt.Println("delete err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 5 -> 7

	ll.DisplayBackward()
	fmt.Println()
	// -> 7 -> 5
}
