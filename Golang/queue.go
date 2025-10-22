package main

import "fmt"

type QueueNode struct {
	value string
	next  *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

var queue Queue

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) QPUSH(value string) {
	newNode := &QueueNode{value: value}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) QPOP() string {
	if q.head == nil {
		fmt.Println("Очередь пустая")
		return ""
	}

	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return value
}

func (q *Queue) Print() {
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
