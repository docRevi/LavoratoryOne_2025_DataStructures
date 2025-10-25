package main

import "fmt"

type ListNode[T comparable] struct {
	value T
	next  *ListNode[T]
}

type LinkedList[T comparable] struct {
	head *ListNode[T]
	tail *ListNode[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) LPUSHBack(value T) {
	newNode := &ListNode[T]{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList[T]) LPUSHFront(value T) {
	newNode := &ListNode[T]{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList[T]) LINSERTAfter(value T, index int) {
	if l.head == nil || index < 1 {
		return
	}

	iter := l.head
	for i := 1; i < index; i++ {
		if iter == nil {
			return
		}
		iter = iter.next
	}

	if iter == nil {
		return
	}

	newNode := &ListNode[T]{value: value}
	newNode.next = iter.next
	iter.next = newNode

	if iter == l.tail {
		l.tail = newNode
	}
}

func (l *LinkedList[T]) LINSERTBefore(value T, index int) {
	if l.head == nil || index < 1 {
		return
	}

	if index == 1 {
		l.LPUSHFront(value)
		return
	}

	iter := l.head
	for i := 1; i < index-1; i++ {
		if iter == nil {
			return
		}
		iter = iter.next
	}

	if iter == nil {
		return
	}

	newNode := &ListNode[T]{value: value}
	newNode.next = iter.next
	iter.next = newNode
}

func (l *LinkedList[T]) LDELHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
}

func (l *LinkedList[T]) LDELTail() {
	if l.head == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	iter := l.head
	for iter.next != l.tail {
		iter = iter.next
	}

	iter.next = nil
	l.tail = iter
}

func (l *LinkedList[T]) LDELAfter(index int) {
	if l.head == nil || index < 1 {
		return
	}

	iter := l.head
	for i := 1; i < index; i++ {
		if iter == nil {
			return
		}
		iter = iter.next
	}

	if iter == nil || iter.next == nil {
		return
	}

	iter.next = iter.next.next
	if iter.next == nil {
		l.tail = iter
	}
}

func (l *LinkedList[T]) LDELBefore(index int) {
	if l.head == nil || index <= 1 {
		return
	}

	if index == 2 {
		l.LDELHead()
		return
	}

	iter := l.head
	for i := 1; i < index-2; i++ {
		if iter == nil {
			return
		}
		iter = iter.next
	}

	if iter == nil || iter.next == nil {
		return
	}

	iter.next = iter.next.next
}

func (l *LinkedList[T]) LDEL(value T) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.LDELHead()
		return
	}

	iter := l.head
	for iter.next != nil && iter.next.value != value {
		iter = iter.next
	}

	if iter.next == nil {
		return
	}

	iter.next = iter.next.next
	if iter.next == nil {
		l.tail = iter
	}
}

func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func (l *LinkedList[T]) LGET(value T) int {
	index := 1
	current := l.head

	for current != nil {
		if current.value == value {
			return index
		}
		current = current.next
		index++
	}
	return -1
}
