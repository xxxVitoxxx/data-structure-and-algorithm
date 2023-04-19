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
		ll.head, ll.tail = node, node
		return
	}

	ll.head.previous = node
	node.next = ll.head
	ll.head = node
}

// InsertLast add an element at the end of the list
func (ll *LinkedList) InsertLast(element int) {
	node := &Node{data: element}
	if ll.tail == nil {
		ll.head, ll.tail = node, node
		return
	}

	ll.tail.next = node
	node.previous = ll.tail
	ll.tail = node
}

// InsertAfter add an element after an item of the list
func (ll *LinkedList) InsertAfter(prev, element int) error {
	node := &Node{data: element}
	if ll.head == nil {
		return errors.New("linked list is empty")
	}

	for current := ll.head; current != nil; current = current.next {
		if current.data == prev {
			if current == ll.tail {
				ll.tail.next = node
				node.previous = ll.tail
				ll.tail = node
				return nil
			}

			current.next.previous = node
			node.next = current.next
			current.next = node
			node.previous = current
			return nil
		}
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
	ll.head.previous = nil
	return nil
}

// DeleteLast delete an element from the end of the list
func (ll *LinkedList) DeleteLast() error {
	if ll.tail == nil {
		return errors.New("linked list is empty")
	}

	if ll.head == ll.tail {
		ll.head, ll.tail = nil, nil
		return nil
	}

	ll.tail = ll.tail.previous
	ll.tail.next = nil
	return nil
}

// Delete delete an element from the list using the key
func (ll *LinkedList) Delete(key int) error {
	if ll.head == nil {
		return errors.New("linked list is empty")
	}

	for current := ll.head; current != nil; current = current.next {
		if current.data == key {
			if current == ll.head {
				ll.head = ll.head.next
				ll.head.previous = nil
				return nil
			}

			if current == ll.tail {
				ll.tail = ll.tail.previous
				ll.tail.next = nil
				return nil
			}

			current.next.previous = current.previous
			current.previous.next = current.next
			return nil
		}
	}
	return errors.New("key not found")
}

// DisplayForward display the complete list in a forward manner
func (ll *LinkedList) DisplayForward() {
	for current := ll.head; current != nil; current = current.next {
		fmt.Printf(" -> %v", current.data)
	}
}

// DisplayBackward display the complete list in a backward manner
func (ll *LinkedList) DisplayBackward() {
	for current := ll.tail; current != nil; current = current.previous {
		fmt.Printf(" -> %v", current.data)
	}
}

func main() {
	ll := NewLinkedList()
	fmt.Println("insertion")
	ll.Insertion(3)
	ll.Insertion(8)
	ll.Insertion(1)

	ll.DisplayForward()
	fmt.Println()
	// -> 1 -> 8 -> 3

	ll.DisplayBackward()
	fmt.Println()
	// -> 3 -> 8 -> 1

	fmt.Println()
	fmt.Println("insert last")
	ll.InsertLast(5)

	ll.DisplayForward()
	fmt.Println()
	// -> 1 -> 8 -> 3 -> 5

	ll.DisplayBackward()
	fmt.Println()
	// -> 5 -> 3 -> 8 -> 1

	fmt.Println()
	fmt.Println("insert after")
	err := ll.InsertAfter(8, 9)
	if err != nil {
		fmt.Println("insert after err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 1 -> 8 -> 9 -> 3 -> 5

	ll.DisplayBackward()
	fmt.Println()
	// -> 5 -> 3 -> 9 -> 8 -> 1

	fmt.Println()
	fmt.Println("deletion")
	err = ll.Deletion()
	if err != nil {
		fmt.Println("deletion err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 8 -> 9 -> 3 -> 5

	ll.DisplayBackward()
	fmt.Println()
	// -> 5 -> 3 -> 9 -> 8

	fmt.Println()
	fmt.Println("delete after")
	err = ll.DeleteLast()
	if err != nil {
		fmt.Println("delete last err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 8 -> 9 -> 3

	ll.DisplayBackward()
	fmt.Println()
	// -> 3 -> 9 -> 8

	fmt.Println()
	fmt.Println("delete")

	err = ll.Delete(9)
	if err != nil {
		fmt.Println("delete err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 8 -> 3

	ll.DisplayBackward()
	fmt.Println()
	// -> 3 -> 8
}
