package main

import "fmt"

func main() {
	readers := map[string]map[string]string{
		"Петров": {
			"Война и мир":        "Лев Толстой",
			"Мастер и Маргарита": "Михаил Булгаков",
		},
		"Сидоров": {
			"Подземный кароль":         "Федор Достоевский",
			"Преступление и наказание": "Федор Достоевский",
			"Вавилон":                  "Гай Ричи",
		},
	}
	fmt.Println(len(readers)) // количество пользователей
	for name, books := range readers {
		fmt.Println(name, "на руках книг: ", len(books))
		for book, author := range books {
			fmt.Printf("\t%v - %v\n", book, author)
		}
	}
}
