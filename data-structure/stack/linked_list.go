package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) Peek() int {
	return ll.head.data
}

func (ll *LinkedList) Print() {
	for current := ll.head; current != nil; current = current.next {
		fmt.Printf("%d  ", current.data)
	}
	fmt.Println()
}

func (ll *LinkedList) Push(data int) {
	ll.head = &Node{
		data: data,
		next: ll.head,
	}
	ll.size++
}

func (ll *LinkedList) Pop() {
	ll.head = ll.head.next
	ll.size--
}

func main() {
	linkedList := NewLinkedList()
	linkedList.Push(1)
	linkedList.Push(2)
	linkedList.Print()
	// 2  1

	fmt.Println("size: ", linkedList.Size())
	// size:  2

	linkedList.Push(3)
	fmt.Println("peek: ", linkedList.Peek())
	// peek:  3

	linkedList.Pop()
	linkedList.Print()
	// 2  1
}
