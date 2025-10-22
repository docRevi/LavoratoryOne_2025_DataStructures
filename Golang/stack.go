package main

import "fmt"

type StackNode struct {
	value string
	next  *StackNode
}

type Stack struct {
	top *StackNode
}

var stack Stack

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) SPUSH(value string) {
	newNode := &StackNode{value: value}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) SPOP() string {
	if s.top == nil {
		fmt.Println("Стек пустой")
		return ""
	}

	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *Stack) Print() {
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
