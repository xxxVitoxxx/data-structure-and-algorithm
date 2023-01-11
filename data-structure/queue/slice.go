package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	queue    []int
	capacity int
}

func NewQueue(cap int) *Queue {
	return &Queue{capacity: cap}
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue) IsFull() bool {
	return len(q.queue) == q.capacity
}

// Enqueue add an element to the front of the queue
func (q *Queue) Enqueue(element int) error {
	if q.IsFull() {
		return errors.New("queue overflow")
	}

	q.queue = append(q.queue, element)
	return nil
}

// Dequeue remove the first element from a queue
func (q *Queue) Dequeue() error {
	if q.IsEmpty() {
		return errors.New("queue underflow")
	}

	q.queue = q.queue[1:]
	return nil
}

// Peek return the first element from our queue without updating queue
func (q *Queue) Peek() (int, error) {
	if len(q.queue) == 0 {
		return 0, errors.New("empty queue")
	}
	return q.queue[0], nil
}

func main() {
	queue := NewQueue(10)
	queue.Enqueue(1)
	queue.Enqueue(10)
	peek, _ := queue.Peek()
	fmt.Println(peek)
	// 1

	queue.Dequeue()
	peek, _ = queue.Peek()
	fmt.Println(peek)
	// 10
}
