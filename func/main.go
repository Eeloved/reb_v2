package main

import (
	"fmt"
)

type wordLogger func(w ...string) (cnt int)

func contains(a []string, x string) bool {
	for _, item := range a {
		if item == x {
			return true
		}
	}
	return false
}

func getMax(nums ...int) int {
	if len(nums) == 0 {
		panic("getMax requires at least one argument")
	}

	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	words := []string{"apple", "banana", "cherry"}
	fmt.Println(contains(words, "banana")) // true
	fmt.Println(contains(words, "grape"))  // false

	fmt.Println(getMax(10, 20, 5, 7, 30)) // 30
	fmt.Println(getMax(-1, -5, -3))       // -1
}

/*func main() {
	customer := map[string]string{
		"name": "John",
		"lastName": "Smith",
	}

			if lastName, isExist := customer["lastName"]; isExist {
				fmt.Println("Hello Mr.", lastName)
			}
		if _, isExist := customer["lastName"]; !isExist {
			fmt.Println("Empty field!")
		}

	userRole := "guest"
	switch userRole {
	case "admin", "user":
		fmt.Println("Access granted!")
	case "guest":
		fallthrough
	default:
		fmt.Println("Access denied!")
	}

	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	for _, v := range m {
		fmt.Println(v)
	}
	printWord := func(w ...string) (cnt int) {
		for _, word := range w {
			fmt.Println(word)

		}
		return len(w)
	}
	func(printer wordLogger, w ...string) {
		fmt.Printf("Count words: %d\n", printer(w...))
	}(printWord, "str1", "str2")
}*/
