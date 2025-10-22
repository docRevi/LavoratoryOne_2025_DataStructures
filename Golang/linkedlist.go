package main

import "fmt"

type ListNode struct {
	value string
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

var list LinkedList

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) LPUSHBack(value string) {
	newNode := &ListNode{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) LPUSHFront(value string) {
	newNode := &ListNode{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) LINSERTAfter(value string, index int) {
	if l.head == nil {
		return
	}

	iter := l.head
	for i := 1; i < index && iter != nil; i++ {
		iter = iter.next
	}
	if iter == nil {
		return
	}

	newNode := &ListNode{value: value}
	newNode.next = iter.next
	iter.next = newNode

	if iter == l.tail {
		l.tail = newNode
	}
}

func (l *LinkedList) LINSERTBefore(value string, index int) {
	if l.head == nil {
		return
	}

	if index == 1 {
		l.LPUSHFront(value)
		return
	}

	iter := l.head
	for i := 1; i < index-1 && iter != nil; i++ {
		iter = iter.next
	}
	if iter == nil {
		return
	}

	newNode := &ListNode{value: value}
	newNode.next = iter.next
	iter.next = newNode
}

func (l *LinkedList) LDELHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
}

func (l *LinkedList) LDELTail() {
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

func (l *LinkedList) LDELAfter(index int) {
	if l.head == nil {
		return
	}

	iter := l.head
	for i := 1; i < index && iter != nil; i++ {
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

func (l *LinkedList) LDELBefore(index int) {
	if l.head == nil || index <= 1 {
		return
	}

	if index == 2 {
		l.LDELHead()
		return
	}

	iter := l.head
	for i := 1; i < index-2 && iter != nil; i++ {
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

func (l *LinkedList) LDEL(value string) {
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

	if iter.next != nil {
		iter.next = iter.next.next
		if iter.next == nil {
			l.tail = iter
		}
	}
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func (l *LinkedList) LGET(value string) int {
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
