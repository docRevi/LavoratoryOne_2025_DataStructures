package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	arr   = NewArray[string]()
	list  = NewLinkedList[string]()
	DList = NewDoubleLinkedList[string]()
	stack = NewStack[string]()
	queue = NewQueue[string]()
	tree  = NewTree()
)

func save_array(filename string, arr *Array[string]) {
	file, _ := os.Create(filename)
	defer file.Close()

	for i := 0; i < arr.size; i++ {
		file.WriteString(arr.data[i] + " ")
	}
}

func load_array(filename string, arr *Array[string]) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		value := scanner.Text()
		if value != "" {
			arr.Apush_back(value)
		}
	}
}

func save_list(filename string, list *LinkedList[string]) {
	file, _ := os.Create(filename)
	defer file.Close()

	current := list.head
	for current != nil {
		file.WriteString(current.value + " ")
		current = current.next
	}
}

func load_list(filename string, list *LinkedList[string]) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		value := scanner.Text()
		if value != "" {
			list.LPUSHBack(value)
		}
	}
}

func save_dlist(filename string, dlist *DoubleLinkedList[string]) {
	file, _ := os.Create(filename)
	defer file.Close()

	iter := dlist.head
	for iter != nil {
		file.WriteString(iter.value + " ")
		iter = iter.next
	}
}

func load_dlist(filename string, dlist *DoubleLinkedList[string]) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		value := scanner.Text()
		if value != "" {
			dlist.DPUSHBack(value)
		}
	}
}

func save_queue(filename string, q *Queue[string]) {
	file, _ := os.Create(filename)
	defer file.Close()

	iter := q.head
	for iter != nil {
		file.WriteString(iter.value + " ")
		iter = iter.next
	}
}

func load_queue(filename string, q *Queue[string]) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		value := scanner.Text()
		if value != "" {
			q.QPUSH(value)
		}
	}
}

func save_stack(filename string, s *Stack[string]) {
	file, _ := os.Create(filename)
	defer file.Close()

	iter := s.top
	for iter != nil {
		file.WriteString(iter.value + " ")
		iter = iter.next
	}
}

func load_stack(filename string, s *Stack[string]) {
	file, _ := os.Open(filename)
	defer file.Close()

	var values []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		value := scanner.Text()
		if value != "" {
			values = append(values, value)
		}
	}

	for i := len(values) - 1; i >= 0; i-- {
		s.SPUSH(values[i])
	}
}

func save_tree_helper(node *TreeNode, file *os.File) {
	if node == nil {
		return
	}
	file.WriteString(strconv.Itoa(node.value) + " ")
	save_tree_helper(node.left, file)
	save_tree_helper(node.right, file)
}

func save_tree(filename string, t *Tree) {
	file, _ := os.Create(filename)
	defer file.Close()

	save_tree_helper(t.root, file)
}

func load_tree(filename string, t *Tree) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		valueStr := scanner.Text()
		if valueStr != "" {
			value, _ := strconv.Atoi(valueStr)
			t.TPUSH(value)
		}
	}
}

func load_all(filename string) {
	load_array(filename+".array", arr)
	load_list(filename+".list", list)
	load_dlist(filename+".dlist", DList)
	load_queue(filename+".queue", queue)
	load_stack(filename+".stack", stack)
	load_tree(filename+".tree", tree)
}

func save_all(filename string) {
	save_array(filename+".array", arr)
	save_list(filename+".list", list)
	save_dlist(filename+".dblist", DList)
	save_queue(filename+".queue", queue)
	save_stack(filename+".stack", stack)
	save_tree(filename+".tree", tree)
}

func execute_query(query string) {
	parts := strings.Fields(query)
	if len(parts) == 0 {
		return
	}

	cmd := parts[0]

	switch cmd {
	case "APUSHB":
		if len(parts) > 1 {
			arr.Apush_back(parts[1])
		}
	case "ADEL":
		if len(parts) > 1 {
			idx, _ := strconv.Atoi(parts[1])
			arr.Adel(idx)
		}
	case "AGET":
		if len(parts) > 1 {
			idx, _ := strconv.Atoi(parts[1])
			result := arr.Aget(idx)
			if result != "" {
				fmt.Println(result)
			} else {
				fmt.Println("Индекс вне диапазона")
			}
		}
	case "ASIZE":
		fmt.Println("Размер массива:", arr.Asize())
	case "AADD":
		if len(parts) > 2 {
			idx, _ := strconv.Atoi(parts[1])
			arr.Aadd(idx, parts[2])
		}
	case "AREPLACE":
		if len(parts) > 2 {
			idx, _ := strconv.Atoi(parts[1])
			arr.Areplace(idx, parts[2])
		}
	case "LPUSHF":
		if len(parts) > 1 {
			list.LPUSHFront(parts[1])
		}
	case "LPUSHB":
		if len(parts) > 1 {
			list.LPUSHBack(parts[1])
		}
	case "LDEL":
		if len(parts) > 1 {
			list.LDEL(parts[1])
		}
	case "LDELH":
		list.LDELHead()
	case "LDELT":
		list.LDELTail()
	case "LDELA":
		if len(parts) > 1 {
			idx, _ := strconv.Atoi(parts[1])
			list.LDELAfter(idx)
		}
	case "LDELB":
		if len(parts) > 1 {
			idx, _ := strconv.Atoi(parts[1])
			list.LDELBefore(idx)
		}
	case "LGET":
		if len(parts) > 1 {
			result := list.LGET(parts[1])
			if result != -1 {
				fmt.Println("Найдено на позиции:", result)
			} else {
				fmt.Println("Элемент не найден")
			}
		}
	case "LINSERTA":
		if len(parts) > 2 {
			idx, _ := strconv.Atoi(parts[1])
			list.LINSERTAfter(parts[2], idx)
		}
	case "LINSERTB":
		if len(parts) > 2 {
			idx, _ := strconv.Atoi(parts[1])
			list.LINSERTBefore(parts[2], idx)
		}
	case "DPUSHF":
		if len(parts) > 1 {
			DList.DPUSHFront(parts[1])
		}
	case "DPUSHB":
		if len(parts) > 1 {
			DList.DPUSHBack(parts[1])
		}
	case "DDEL":
		if len(parts) > 1 {
			DList.DDELValue(parts[1])
		}
	case "DDELH":
		DList.DDELHead()
	case "DDELT":
		DList.DDELTail()
	case "DDELA":
		if len(parts) > 1 {
			idx, _ := strconv.Atoi(parts[1])
			DList.DDELAfter(idx)
		}
	case "DDELB":
		if len(parts) > 1 {
			idx, _ := strconv.Atoi(parts[1])
			DList.DDELBefore(idx)
		}
	case "DGET":
		if len(parts) > 1 {
			result := DList.DGET(parts[1])
			if result != -1 {
				fmt.Println("Найдено на позиции:", result)
			} else {
				fmt.Println("Элемент не найден")
			}
		}
	case "DINSERTA":
		if len(parts) > 2 {
			idx, _ := strconv.Atoi(parts[1])
			DList.DINSERTAfter(parts[2], idx)
		}
	case "DINSERTB":
		if len(parts) > 2 {
			idx, _ := strconv.Atoi(parts[1])
			DList.DINSERTBefore(parts[2], idx)
		}
	case "QPUSH":
		if len(parts) > 1 {
			queue.QPUSH(parts[1])
		}
	case "QPOP":
		result := queue.QPOP()
		if result != "" {
			fmt.Println("Извлечено:", result)
		}
	case "SPUSH":
		if len(parts) > 1 {
			stack.SPUSH(parts[1])
		}
	case "SPOP":
		result := stack.SPOP()
		if result != "" {
			fmt.Println("Извлечено:", result)
		}
	case "TPUSH":
		if len(parts) > 1 {
			val, _ := strconv.Atoi(parts[1])
			tree.TPUSH(val)
		}
	case "TSEARCH":
		if len(parts) > 1 {
			val, _ := strconv.Atoi(parts[1])
			found := tree.TSEARCH(tree.root, val)
			if found {
				fmt.Println("Найдено")
			} else {
				fmt.Println("Не найдено")
			}
		}
	case "ISFULL":
		isFull := tree.isFullBinaryTree(tree.root)
		if isFull {
			fmt.Println("Полное бинарное дерево")
		} else {
			fmt.Println("Не полное бинарное дерево")
		}
	case "PRINT":
		if len(parts) > 1 {
			switch parts[1] {
			case "list":
				list.Print()
			case "dlist":
				DList.PrintForward()
			case "array":
				arr.Aprint()
			case "queue":
				queue.Print()
			case "stack":
				stack.Print()
			case "tree":
				tree.print(tree.root)
			default:
				fmt.Println("Неизвестная структура:", parts[1])
			}
		}
	default:
		fmt.Println("Неизвестная команда:", cmd)
	}
}
