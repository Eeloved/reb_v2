package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inFilePath := filepath.Join("data", "in.txt")
	outFilePath := filepath.Join("data", "data_out.txt")

	in, err := os.Open(inFilePath)
	if err != nil {
		panic(fmt.Sprintf("cannot open input file: %v", err))
	}
	defer in.Close()

	out, err := os.Create(outFilePath)
	if err != nil {
		panic(fmt.Sprintf("cannot create output file: %v", err))
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	defer func() {
		if r := recover(); r != nil {
			writer.Flush()
			fmt.Println("Recovered from panic:", r)

			// Read and print contents written before panic
			content, err := os.ReadFile(outFilePath)
			if err != nil {
				fmt.Printf("Could not read output file: %v\n", err)
				return
			}
			fmt.Println("\nContent written before panic:\n")
			fmt.Println(string(content))
		}
	}()

	scanner := bufio.NewScanner(in)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")

		if len(fields) != 3 || strings.TrimSpace(fields[0]) == "" || strings.TrimSpace(fields[1]) == "" || strings.TrimSpace(fields[2]) == "" {
			panic(fmt.Sprintf("parse error: empty field on string %d", lineNum))
		}

		output := fmt.Sprintf("Row: %d\nName : %s\nAddress: %s\nCity: %s\n\n\n",
			lineNum,
			strings.TrimSpace(fields[0]),
			strings.TrimSpace(fields[1]),
			strings.TrimSpace(fields[2]),
		)

		_, err := writer.WriteString(output)
		if err != nil {
			panic(fmt.Sprintf("write error: %v", err))
		}

		lineNum++
	}

	writer.Flush()

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("scanner error: %v", err))
	}
}

/*package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	inFilePath := filepath.Join("data", "in.txt")
	outFilePath := filepath.Join("data", "out.txt")
	// Открываем входной файл
	inFile, err := os.Open(inFilePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer inFile.Close()

	// Создаем (или очищаем) выходной файл
	outFile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println("Ошибка при создании выходного файла:", err)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)

	row := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Разделяем строку по символу '|'
		parts := strings.SplitN(line, "|", 3)
		if len(parts) != 3 ||
			strings.TrimSpace(parts[0]) == "" ||
			strings.TrimSpace(parts[1]) == "" ||
			strings.TrimSpace(parts[2]) == "" {
			fmt.Printf("Пропущена строка %d: неполные данные\n", row)
			continue
		}

		name := strings.TrimSpace(parts[0])
		address := strings.TrimSpace(parts[1])
		city := strings.TrimSpace(parts[2])

		// Пишем в файл в нужном формате
		fmt.Fprintf(writer, "Row: %d\nName : %s\nAddress: %s\nCity: %s\n\n\n", row, name, address, city)
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}

	writer.Flush()
}
*/
