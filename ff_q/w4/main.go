package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func mapString() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   17,
		"Eve":   33,
	}

	ages["Charlie"] = 19 // add
	age := ages["Alice"]

	//age, exists := ages["Alice"]

	delete(ages, "Bob")

	fmt.Println(age, ages)
}

func stringsFunc() {

	words := strings.Split("hello world go", " ")
	fmt.Println(words)

	count := len(words)
	fmt.Println(count)

	joined := strings.Join(words, "-")
	fmt.Println(joined)

	num, _ := strconv.Atoi("42") // string to int

	fmt.Println(num)

	str := strconv.Itoa(42) // int to string
	fmt.Println(str)

	f, _ := strconv.ParseFloat("3.14", 64) // string to float
	fmt.Println(f)
}

func userStruct() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{Name: "Alice", Age: 25}
	fmt.Println(user)
	data, _ := json.Marshal(user)

	fmt.Println(string(data))

}

func main() {
	//mapString()
	//stringsFunc()
	userStruct()
}
