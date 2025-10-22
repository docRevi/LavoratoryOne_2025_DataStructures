package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ExecuteQuery(query string) {
	parts := strings.Fields(query)
	if len(parts) == 0 {
		return
	}

	cmd := parts[0]

	switch cmd {
	// Команды массива
	case "APUSHB":
		if len(parts) > 1 {
			arr.Apush_back(parts[1])
		}
	case "ADEL":
		if len(parts) > 1 {
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				arr.Adel(idx)
			}
		}
	case "AGET":
		if len(parts) > 1 {
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				result := arr.Aget(idx)
				if result != "" {
					fmt.Println(result)
				} else {
					fmt.Println("Индекс вне диапазона")
				}
			}
		}
	case "ASIZE":
		fmt.Println("Размер массива:", arr.Asize())
	case "AADD":
		if len(parts) > 2 {
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				arr.Aadd(idx, parts[2])
			}
		}
	case "AREPLACE":
		if len(parts) > 2 {
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				arr.Areplace(idx, parts[2])
			}
		}

	// Команды односвязного списка
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
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				list.LDELAfter(idx)
			}
		}
	case "LDELB":
		if len(parts) > 1 {
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				list.LDELBefore(idx)
			}
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
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				list.LINSERTAfter(parts[2], idx)
			}
		}
	case "LINSERTB":
		if len(parts) > 2 {
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				list.LINSERTBefore(parts[2], idx)
			}
		}

	// Команды двусвязного списка
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
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				DList.DDELAfter(idx)
			}
		}
	case "DDELB":
		if len(parts) > 1 {
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				DList.DDELBefore(idx)
			}
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
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				DList.DINSERTAfter(parts[2], idx)
			}
		}
	case "DINSERTB":
		if len(parts) > 2 {
			if idx, err := strconv.Atoi(parts[1]); err == nil {
				DList.DINSERTBefore(parts[2], idx)
			}
		}

	// Команды очереди
	case "QPUSH":
		if len(parts) > 1 {
			queue.QPUSH(parts[1])
		}
	case "QPOP":
		result := queue.QPOP()
		if result != "" {
			fmt.Println("Извлечено:", result)
		}

	// Команды стека
	case "SPUSH":
		if len(parts) > 1 {
			stack.SPUSH(parts[1])
		}
	case "SPOP":
		result := stack.SPOP()
		if result != "" {
			fmt.Println("Извлечено:", result)
		}

	// Команды дерева
	case "TPUSH":
		if len(parts) > 1 {
			if value, err := strconv.Atoi(parts[1]); err == nil {
				tree.TPUSH(value)
			}
		}
	case "TSEARCH":
		if len(parts) > 1 {
			if value, err := strconv.Atoi(parts[1]); err == nil {
				found := tree.TSEARCH(value)
				if found {
					fmt.Println("Найдено")
				} else {
					fmt.Println("Не найдено")
				}
			}
		}
	case "ISFULL":
		isFull := tree.IsFullBinaryTree()
		if isFull {
			fmt.Println("Полное бинарное дерево")
		} else {
			fmt.Println("Не полное бинарное дерево")
		}

	// Команды вывода
	case "PRINT":
		if len(parts) > 1 {
			target := parts[1]
			switch target {
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
				tree.Print()
			default:
				fmt.Println("Неизвестная структура:", target)
			}
		}

	default:
		fmt.Println("Неизвестная команда:", cmd)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println("Односвязный список: LPUSHF, LPUSHB, LDEL, LDELH, LDELT, LDELA, LDELB, LGET, LINSERTA, LINSERTB")
	fmt.Println("Двусвязный список: DPUSHF, DPUSHB, DDEL, DDELH, DDELT, DDELA, DDELB, DGET, DINSERTA, DINSERTB")
	fmt.Println("Массив: APUSHB, ADEL, AGET, ASIZE, AADD, AREPLACE")
	fmt.Println("Очередь: QPUSH, QPOP")
	fmt.Println("Стек: SPUSH, SPOP")
	fmt.Println("Дерево: TPUSH, TSEARCH, ISFULL")
	fmt.Println("Вывод: PRINT <структура>")
}
