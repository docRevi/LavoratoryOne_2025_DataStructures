package main

import "fmt"

type Array[T comparable] struct {
	data     []T
	size     int
	capacity int
}

func NewArray[T comparable]() *Array[T] {
	return &Array[T]{
		data:     make([]T, 0),
		size:     0,
		capacity: 0,
	}
}

func (a *Array[T]) Reserve(newCapacity int) {
	if newCapacity <= a.capacity {
		return
	}
	newData := make([]T, newCapacity)
	copy(newData, a.data)
	a.data = newData
	a.capacity = newCapacity
}

func (a *Array[T]) Aadd(index int, value T) {
	if index < 0 || index > a.size {
		return
	}
	if a.size == a.capacity {
		a.Reserve(a.capacity*2 + 1)
	}

	a.data = append(a.data, value)
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = value
	a.size++
}

func (a *Array[T]) Apush_back(value T) {
	a.Aadd(a.size, value)
}

func (a *Array[T]) Aget(index int) T {
	var zero T
	if index < 0 || index >= a.size {
		return zero
	}
	return a.data[index]
}

func (a *Array[T]) Adel(index int) {
	if index < 0 || index >= a.size {
		return
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	a.size--
}

func (a *Array[T]) Areplace(index int, value T) {
	if index < 0 || index >= a.size {
		return
	}
	a.data[index] = value
}

func (a *Array[T]) Asize() int {
	return a.size
}

func (a *Array[T]) Aprint() {
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i], " ")
	}
	fmt.Println()
}
