package main

import (
	"bufio"
	"fmt"
	"os"
)

var filename string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите файл с помощью --file <имя_файла>")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--file" && i+1 < len(os.Args) {
			filename = os.Args[i+1]
			i++
		} else if arg == "--query" && i+1 < len(os.Args) {
			query := os.Args[i+1]
			ExecuteQuery(query)
			SaveAll()
			return
		}
	}

	if filename == "" {
		fmt.Println("Укажите файл с помощью --file <имя_файла>")
		return
	}

	LoadAll()

	fmt.Println("Введите команду (или exit):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}
		ExecuteQuery(line)
		SaveAll()
	}
}

func LoadAll() {
	LoadArray(filename + ".array")
	LoadList(filename + ".list")
	LoadDList(filename + ".dlist")
	LoadQueue(filename + ".queue")
	LoadStack(filename + ".stack")
	LoadTree(filename + ".tree")
}

func SaveAll() {
	SaveArray(filename + ".array")
	SaveList(filename + ".list")
	SaveDList(filename + ".dblist")
	SaveQueue(filename + ".queue")
	SaveStack(filename + ".stack")
	SaveTree(filename + ".tree")
}
