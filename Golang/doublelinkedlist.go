package main

import "fmt"

type DListNode struct {
	value string
	next  *DListNode
	prev  *DListNode
}

type DoubleLinkedList struct {
	head *DListNode
	tail *DListNode
}

var DList DoubleLinkedList

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (d *DoubleLinkedList) DPUSHFront(value string) {
	newNode := &DListNode{value: value}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}
	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode
}

func (d *DoubleLinkedList) DPUSHBack(value string) {
	newNode := &DListNode{value: value}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}
	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode
}

func (d *DoubleLinkedList) DINSERTAfter(value string, key int) {
	if d.head == nil || key < 1 {
		return
	}

	iter := d.head
	for i := 1; i < key && iter != nil; i++ {
		iter = iter.next
	}
	if iter == nil {
		return
	}

	newNode := &DListNode{value: value}
	newNode.next = iter.next
	newNode.prev = iter

	if iter.next != nil {
		iter.next.prev = newNode
	} else {
		d.tail = newNode
	}
	iter.next = newNode
}

func (d *DoubleLinkedList) DINSERTBefore(value string, key int) {
	if d.head == nil || key < 1 {
		return
	}

	if key == 1 {
		d.DPUSHFront(value)
		return
	}

	iter := d.head
	for i := 1; i < key && iter != nil; i++ {
		iter = iter.next
	}
	if iter == nil {
		return
	}

	newNode := &DListNode{value: value}
	newNode.prev = iter.prev
	newNode.next = iter
	iter.prev.next = newNode
	iter.prev = newNode
}

func (d *DoubleLinkedList) DDELHead() {
	if d.head == nil {
		return
	}
	if d.head.next == nil {
		d.head = nil
		d.tail = nil
		return
	}
	d.head = d.head.next
	d.head.prev = nil
}

func (d *DoubleLinkedList) DDELTail() {
	if d.head == nil {
		return
	}
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
		return
	}
	d.tail = d.tail.prev
	d.tail.next = nil
}

func (d *DoubleLinkedList) DDELBefore(index int) {
	if d.head == nil || index < 1 {
		return
	}

	if index == 1 {
		d.DDELHead()
		return
	}

	iter := d.head
	for i := 1; i < index && iter != nil; i++ {
		iter = iter.next
	}
	if iter == nil || iter.prev == nil {
		return
	}

	nodeToDelete := iter.prev
	if nodeToDelete.prev != nil {
		nodeToDelete.prev.next = iter
	} else {
		d.head = iter
	}
	iter.prev = nodeToDelete.prev
}

func (d *DoubleLinkedList) DDELAfter(index int) {
	if d.head == nil || index < 1 {
		return
	}

	iter := d.head
	for i := 1; i < index && iter != nil; i++ {
		iter = iter.next
	}
	if iter == nil || iter.next == nil {
		return
	}

	nodeToDelete := iter.next
	iter.next = nodeToDelete.next
	if nodeToDelete.next != nil {
		nodeToDelete.next.prev = iter
	} else {
		d.tail = iter
	}
}

func (d *DoubleLinkedList) DDELValue(value string) {
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

func (d *DoubleLinkedList) DGET(value string) int {
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

func (d *DoubleLinkedList) PrintForward() {
	iter := d.head
	for iter != nil {
		fmt.Print(iter.value, " ")
		iter = iter.next
	}
	fmt.Println()
}

func (d *DoubleLinkedList) PrintBackward() {
	iter := d.tail
	for iter != nil {
		fmt.Print(iter.value, " ")
		iter = iter.prev
	}
	fmt.Println()
}
