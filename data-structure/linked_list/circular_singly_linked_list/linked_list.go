package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
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
		node.next = node
		ll.head, ll.tail = node, node
		return
	}

	node.next = ll.head
	ll.tail.next = node
	ll.head = node
}

// InsertLast add an element at the end of the list
func (ll *LinkedList) InsertLast(element int) {
	node := &Node{data: element}
	if ll.tail == nil {
		node.next = node
		ll.head, ll.tail = node, node
		return
	}

	ll.tail.next = node
	node.next = ll.head
	ll.tail = node
}

// InsertAfter add an element after an item of the list
func (ll *LinkedList) InsertAfter(prev, element int) error {
	if ll.head == nil {
		return errors.New("linked list is empty")
	}

	node := &Node{data: element}
	current := ll.head
	for {
		if current.data == prev {
			if current == ll.tail {
				ll.tail.next = node
				node.next = ll.head
				ll.tail = node
				return nil
			}

			node.next = current.next
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

	ll.tail.next = ll.head.next
	ll.head = ll.head.next
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

	var current *Node
	for current = ll.head; current.next != ll.tail; current = current.next {
	}

	current.next = ll.head
	ll.tail = current
	return nil
}

// Delete delete an element from the list using the key
func (ll *LinkedList) Delete(key int) error {
	if ll.head == nil {
		return errors.New("linked list is empty")
	}

	if ll.head.data == key {
		if ll.head == ll.tail {
			ll.head, ll.tail = nil, nil
			return nil
		}

		ll.tail.next = ll.head.next
		ll.head = ll.head.next
		return nil
	}

	prev := &Node{}
	current := ll.head
	for current != ll.tail {
		prev = current
		current = current.next
		if current.data == key {
			if current == ll.tail {
				prev.next = current.next
				ll.tail = prev
				return nil
			}

			prev.next = current.next
			return nil
		}
	}
	return errors.New("key not found")
}

// DisplayForward display the complete list in a forward manner
func (ll *LinkedList) DisplayForward() {
	current := ll.head
	if current == nil {
		return
	}

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
	current := ll.head
	if current == nil {
		return
	}

	var stack []int
	for {
		stack = append(stack, current.data)
		if current == ll.tail {
			break
		}
		current = current.next
	}

	for len(stack) != 0 {
		fmt.Printf(" -> %d", stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
}

func main() {
	ll := NewLinkedList()
	fmt.Println("insertion")
	ll.Insertion(6)
	ll.Insertion(9)
	ll.Insertion(2)

	ll.DisplayForward()
	fmt.Println()
	// -> 2 -> 9 -> 6

	ll.DisplayBackward()
	fmt.Println()
	// -> 6 -> 9 -> 2

	fmt.Println()
	fmt.Println("insert last")
	ll.InsertLast(7)

	ll.DisplayForward()
	fmt.Println()
	// -> 2 -> 9 -> 6 ->7

	ll.DisplayBackward()
	fmt.Println()
	// -> 7 -> 6 -> 9 -> 2

	fmt.Println()
	fmt.Println("insert after")
	err := ll.InsertAfter(6, 8)
	if err != nil {
		fmt.Println("insert after err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 2 -> 9 -> 6 -> 8 -> 7

	ll.DisplayBackward()
	fmt.Println()
	// -> 7 -> 8 -> 6 -> 9 -> 2

	fmt.Println()
	fmt.Println("deletion")
	err = ll.Deletion()
	if err != nil {
		fmt.Println("deletion err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 9 -> 6 -> 8 -> 7

	ll.DisplayBackward()
	fmt.Println()
	// -> 7 -> 8 -> 6 -> 9

	fmt.Println()
	fmt.Println("delete last")
	err = ll.DeleteLast()
	if err != nil {
		fmt.Println("delete last err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 9 -> 6 -> 8

	ll.DisplayBackward()
	fmt.Println()
	// -> 8 -> 6 -> 9

	fmt.Println()
	fmt.Println("delete")
	err = ll.Delete(9)
	if err != nil {
		fmt.Println("delete err: ", err)
	}

	ll.DisplayForward()
	fmt.Println()
	// -> 6 -> 8

	ll.DisplayBackward()
	fmt.Println()
	// -> 8 -> 6
}
