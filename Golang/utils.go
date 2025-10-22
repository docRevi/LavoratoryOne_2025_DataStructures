package main

import (
	"bufio"
	"os"
	"strconv"
)

// Функции сохранения/загрузки для массива
func SaveArray(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	for i := 0; i < arr.size; i++ {
		file.WriteString(arr.data[i] + "\n")
	}
}

func LoadArray(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			arr.Apush_back(line)
		}
	}
}

// Функции сохранения/загрузки для односвязного списка
func SaveList(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	iter := list.head
	for iter != nil {
		file.WriteString(iter.value + "\n")
		iter = iter.next
	}
}

func LoadList(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			list.LPUSHBack(line)
		}
	}
}

// Функции сохранения/загрузки для двусвязного списка
func SaveDList(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	iter := DList.head
	for iter != nil {
		file.WriteString(iter.value + "\n")
		iter = iter.next
	}
}

func LoadDList(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			DList.DPUSHBack(line)
		}
	}
}

// Функции сохранения/загрузки для очереди
func SaveQueue(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	iter := queue.head
	for iter != nil {
		file.WriteString(iter.value + "\n")
		iter = iter.next
	}
}

func LoadQueue(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			queue.QPUSH(line)
		}
	}
}

// Функции сохранения/загрузки для стека
func SaveStack(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	// Сохраняем в обратном порядке для корректной загрузки
	values := []string{}
	iter := stack.top
	for iter != nil {
		values = append([]string{iter.value}, values...)
		iter = iter.next
	}

	for _, val := range values {
		file.WriteString(val + "\n")
	}
}

func LoadStack(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			stack.SPUSH(line)
		}
	}
}

// Функции сохранения/загрузки для дерева
func saveTreeHelper(node *TreeNode, file *os.File) {
	if node == nil {
		return
	}
	file.WriteString(strconv.Itoa(node.value) + "\n")
	saveTreeHelper(node.left, file)
	saveTreeHelper(node.right, file)
}

func SaveTree(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	saveTreeHelper(tree.root, file)
}

func LoadTree(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if value, err := strconv.Atoi(line); err == nil {
				tree.TPUSH(value)
			}
		}
	}
}
