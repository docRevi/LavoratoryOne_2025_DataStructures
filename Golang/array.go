package main

import (
	"fmt"
)

type Array struct {
	data     []string
	size     int
	capacity int
}

var arr Array

func NewArray() *Array {
	return &Array{
		data:     make([]string, 0),
		size:     0,
		capacity: 0,
	}
}

func (a *Array) Reserve(newCapacity int) {
	if newCapacity <= a.capacity {
		return
	}
	newData := make([]string, newCapacity)
	copy(newData, a.data)
	a.data = newData
	a.capacity = newCapacity
}

func (a *Array) Aadd(index int, value string) {
	if index < 0 || index > a.size {
		return
	}
	if index == a.size {
		a.Reserve(a.size + 1)
		a.size++
	}
	if index < len(a.data) {
		a.data[index] = value
	}
}

func (a *Array) Apush_back(value string) {
	a.data = append(a.data, value)
	a.size = len(a.data)
	a.capacity = cap(a.data)
}

func (a *Array) Aget(index int) string {
	if index < 0 || index >= a.size {
		return ""
	}
	return a.data[index]
}

func (a *Array) Adel(index int) {
	if index < 0 || index >= a.size {
		return
	}
	a.data[index] = ""
}

func (a *Array) Areplace(index int, value string) {
	if index < 0 || index >= a.size {
		return
	}
	a.data[index] = value
}

func (a *Array) Asize() int {
	return a.size
}

func (a *Array) Aprint() {
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i], " ")
	}
	fmt.Println()
}
