package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filename string
	var query string

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "--file" && i+1 < len(args) {
			filename = args[i+1]
			i++
		} else if args[i] == "--query" && i+1 < len(args) {
			query = args[i+1]
			i++
		}
	}

	if filename == "" {
		fmt.Println("Укажите файл с помощью --file <имя_файла>")
		return
	}

	load_all(filename)

	if query != "" {
		execute_query(query)
		save_all(filename)
	} else {
		fmt.Println("Введите команду (или exit):")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "exit" {
				break
			}
			execute_query(line)
			save_all(filename)
		}
	}
}
