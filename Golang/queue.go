package main

import "fmt"

type QueueNode[T comparable] struct {
	value T
	next  *QueueNode[T]
}

type Queue[T comparable] struct {
	head *QueueNode[T]
	tail *QueueNode[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) QPUSH(value T) {
	newNode := &QueueNode[T]{value: value}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue[T]) QPOP() T {
	var zero T
	if q.head == nil {
		fmt.Println("Очередь пустая")
		return zero
	}

	nodeToDelete := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return nodeToDelete.value
}

func (q *Queue[T]) Print() {
	if q.head == nil {
		fmt.Println("Очередь пустая")
		return
	}

	iter := q.head
	for iter != nil {
		fmt.Print(iter.value, " ")
		iter = iter.next
	}
	fmt.Println()
}
