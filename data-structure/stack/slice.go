package main

import (
	"errors"
	"fmt"
)

type StackSlice struct {
	top      int
	capacity int
	stack    []int
}

func NewStackSlice(capacity int) *StackSlice {
	return &StackSlice{
		top:      -1,
		capacity: capacity,
		stack:    make([]int, capacity),
	}
}

func (s *StackSlice) IsEmpty() bool {
	return s.top == -1
}

func (s *StackSlice) IsFull() bool {
	return s.top+1 == s.capacity
}

func (s *StackSlice) Size() int {
	return s.top + 1
}

func (s *StackSlice) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack underflow")
	}
	return s.stack[s.top], nil
}

func (s *StackSlice) Print() {
	for i := 0; i <= s.top; i++ {
		fmt.Printf("%d  ", s.stack[i])
	}
	fmt.Println()
}

func (s *StackSlice) Push(data int) error {
	if s.IsFull() {
		return errors.New("stack overflow")
	}

	s.top++
	s.stack[s.top] = data
	return nil
}

func (s *StackSlice) Pop() error {
	if s.IsEmpty() {
		return errors.New("stack underflow")
	}

	s.stack = s.stack[:s.top]
	s.top--
	return nil
}

func main() {
	stack := NewStackSlice(5)
	stack.Push(1)
	stack.Push(2)
	stack.Print()
	// 1  2

	fmt.Printf("size: %d\n", stack.Size())
	// size: 2

	stack.Push(3)

	peek, _ := stack.Peek()
	fmt.Printf("peek: %d\n", peek)
	// peek: 3

	stack.Pop()
	stack.Print()
	// 1  2
}
