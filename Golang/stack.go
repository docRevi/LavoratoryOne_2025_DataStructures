package main

import "fmt"

type StackNode[T comparable] struct {
	value T
	next  *StackNode[T]
}

type Stack[T comparable] struct {
	top *StackNode[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) SPUSH(value T) {
	newNode := &StackNode[T]{value: value}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack[T]) SPOP() T {
	var zero T
	if s.top == nil {
		fmt.Println("Стек пустой")
		return zero
	}

	nodeToDelete := s.top
	s.top = s.top.next
	return nodeToDelete.value
}

func (s *Stack[T]) Print() {
	if s.top == nil {
		fmt.Println("Стек пустой")
		return
	}

	iter := s.top
	for iter != nil {
		fmt.Print(iter.value, " ")
		iter = iter.next
	}
	fmt.Println()
}
