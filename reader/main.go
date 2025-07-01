package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filePath := filepath.Join("data", "in.txt")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("ошибка закрытия файла: %v", cerr)
		}
	}()

	scanner := bufio.NewScanner(file)
	var count int

	for {
		ok := scanner.Scan()
		if !ok {
			// Проверка на ошибку чтения (включая EOF)
			if err := scanner.Err(); err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("ошибка чтения файла: %v", err)
			}
			break // достигнут EOF без ошибки
		}
		count++
	}

	fmt.Printf("Total strings: %d", count)
}
