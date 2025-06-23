package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var startTime time.Time = time.Now() // фиксируем время старта

// logtime выводит время выполнения программы
func logtime() {
	duration := time.Since(startTime)
	fmt.Printf("Время выполнения программы: %v \n", duration)
}

func main() {
	lineCount, byteCount := 0, 0
	defer logtime() // выполнение при завершении main

	// Пути к файлам
	inFilePath := filepath.Join("data", "in.txt")
	outFilePath := filepath.Join("data", "out.txt")

	// Открытие файла in.txt
	inFile, err := os.Open(inFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка открытия in.txt: %v\n", err)
		return
	}
	defer inFile.Close()

	// Создание или перезапись файла out.txt
	outFile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка создания out.txt: %v\n", err)
		return
	}
	defer func() {
		err := outFile.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при закрытии out.txt: %v\n", err)
			return
		}
		// Печатаем информацию после закрытия out.txt
		fmt.Printf("Всего записано строк: %d, байт: %d\n", lineCount, byteCount)
	}()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// Чтение строк из in.txt и запись в out.txt с нумерацией

	for scanner.Scan() {
		//var lineCount int
		//var byteCount int
		lineCount++
		line := scanner.Text()
		outputLine := fmt.Sprintf("%d: %s\n", lineCount, line)
		n, err := writer.WriteString(outputLine)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка записи в out.txt: %v\n", err)
			return
		}
		byteCount += n
		//fmt.Println(lineCount, byteCount)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения из in.txt: %v\n", err)
		return
	}
}
