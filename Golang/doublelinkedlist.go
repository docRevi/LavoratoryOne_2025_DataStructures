package main

import "fmt"

type DListNode[T comparable] struct {
	value T
	next  *DListNode[T]
	prev  *DListNode[T]
}

type DoubleLinkedList[T comparable] struct {
	head *DListNode[T]
	tail *DListNode[T]
}

func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

func (d *DoubleLinkedList[T]) DPUSHFront(value T) {
	newNode := &DListNode[T]{value: value}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}
	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode
}

func (d *DoubleLinkedList[T]) DPUSHBack(value T) {
	newNode := &DListNode[T]{value: value}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}
	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode
}

func (d *DoubleLinkedList[T]) DINSERTAfter(value T, key int) {
	if d.head == nil || key < 1 {
		return
	}

	iter := d.head
	for i := 1; i < key; i++ {
		if iter == nil {
			return
		}
		iter = iter.next
	}

	if iter == nil {
		return
	}

	newNode := &DListNode[T]{value: value}
	newNode.next = iter.next
	newNode.prev = iter

	if iter.next != nil {
		iter.next.prev = newNode
	} else {
		d.tail = newNode
	}

	iter.next = newNode
}

func (d *DoubleLinkedList[T]) DINSERTBefore(value T, key int) {
	if d.head == nil || key < 1 {
		return
	}

	if key == 1 {
		d.DPUSHFront(value)
		return
	}

	iter := d.head
	for i := 1; i < key-1; i++ {
		if iter == nil {
			return
		}
		iter = iter.next
	}

	if iter == nil {
		return
	}

	newNode := &DListNode[T]{value: value}
	newNode.prev = iter.prev
	newNode.next = iter

	if iter.prev != nil {
		iter.prev.next = newNode
	}
	iter.prev = newNode

	if iter == d.head {
		d.head = newNode
	}
}

func (d *DoubleLinkedList[T]) DDELHead() {
	if d.head == nil {
		return
	}

	if d.head == d.tail {
		d.head = nil
		d.tail = nil
		return
	}

	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	}
}

func (d *DoubleLinkedList[T]) DDELTail() {
	if d.head == nil {
		return
	}

	if d.head == d.tail {
		d.head = nil
		d.tail = nil
		return
	}

	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	}
}

func (d *DoubleLinkedList[T]) DDELBefore(index int) {
	if d.head == nil || index < 1 {
		return
	}

	if index == 1 {
		d.DDELHead()
		return
	}

	iter := d.head
	for i := 1; i < index-1; i++ {
		if iter == nil {
			return
		}
		iter = iter.next
	}

	if iter == nil {
		return
	}

	if iter.prev != nil {
		iter.prev.next = iter.next
	}
	if iter.next != nil {
		iter.next.prev = iter.prev
	}
}

func (d *DoubleLinkedList[T]) DDELAfter(index int) {
	if d.head == nil || index < 1 {
		return
	}

	iter := d.head
	for i := 1; i < index; i++ {
		if iter == nil {
			return
		}
		iter = iter.next
	}

	if iter == nil {
		return
	}

	if iter.prev != nil {
		iter.prev.next = iter.next
	}
	if iter.next != nil {
		iter.next.prev = iter.prev
	}
}

func (d *DoubleLinkedList[T]) DDELValue(value T) {
	if d.head == nil {
		return
	}

	if d.head.value == value {
		d.DDELHead()
		return
	}

	if d.tail.value == value {
		d.DDELTail()
		return
	}

	iter := d.head
	for iter != nil && iter.value != value {
		iter = iter.next
	}

	if iter == nil {
		return
	}

	if iter.prev != nil {
		iter.prev.next = iter.next
	}
	if iter.next != nil {
		iter.next.prev = iter.prev
	}
}

func (d *DoubleLinkedList[T]) DGET(value T) int {
	index := 0
	iter := d.head

	for iter != nil {
		if iter.value == value {
			return index
		}
		iter = iter.next
		index++
	}
	return -1
}

func (d *DoubleLinkedList[T]) PrintForward() {
	iter := d.head
	for iter != nil {
		fmt.Print(iter.value, " ")
		iter = iter.next
	}
	fmt.Println()
}

func (d *DoubleLinkedList[T]) PrintBackward() {
	iter := d.tail
	for iter != nil {
		fmt.Print(iter.value, " ")
		iter = iter.prev
	}
	fmt.Println()
}
