package main

import (
	"errors"
	"fmt"
)

var ErrEmptyQueue error = errors.New("empty queue")

type Node struct {
	data int
	next *Node
}

type Queue struct {
	head, tail *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue add an node to the tail of the linked list
func (q *Queue) Enqueue(element int) {
	node := &Node{data: element}
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

// Dequeue remove the head from a linked list
func (q *Queue) Dequeue() error {
	if q.head == nil {
		return ErrEmptyQueue
	}

	if q.head == q.tail {
		q.head, q.tail = nil, nil
	} else {
		q.head = q.head.next
	}
	return nil
}

// Peek return the first element from our queue without updating queue
func (q *Queue) Peek() (int, error) {
	if q.head == nil {
		return 0, ErrEmptyQueue
	}
	return q.head.data, nil
}

func (q *Queue) Traversal() {
	for current := q.head; current != nil; current = current.next {
		fmt.Printf("-> %d", current.data)
	}
	fmt.Println()
}

func main() {
	q := NewQueue()
	fmt.Println("Enqueue")
	q.Enqueue(10)
	q.Enqueue(7)

	fmt.Println()
	fmt.Println("Peek")
	peek, _ := q.Peek()
	fmt.Println("peek: ", peek)
	// peek:  10

	fmt.Println()
	fmt.Println("Traversal")
	q.Traversal()
	// -> 10-> 7

	fmt.Println()
	fmt.Println("Dequeue")
	q.Dequeue()
	q.Dequeue()

	peek, _ = q.Peek()
	fmt.Println("peek: ", peek)
	// peek: 0
}
