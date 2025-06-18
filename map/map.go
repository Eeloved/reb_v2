package main

import "fmt"

func main() {
	m := map[string]int{"first": 1, "second": 2}
	m["third"] = 3 // добавление

	editMap(m)
	fmt.Println(m)
	//delete(m, "firts")  // удаление

	for key, value := range m {
		fmt.Println(key, value)
	}

	/*			проверка значения
	first, ok := m["second"]
	if !ok {
		fmt.Println("Key not found")
		return
	}
	fmt.Println(first) */
}

func editMap(m map[string]int) {
	m["mew"] = 78
}
