package main

import (
	"fmt"
	"strings"
)
func mapFunc() {

	ages := map[string]int{
		"Alice": 25,
		"Bob":   17,
		"Eve":   30,
		"Charlie": 16,
	}

	for name, age := range ages {
		if age > 18 {
			fmt.Println(name)
		}
	}
}

func WordCount(s string) int {
	
	//s := ("Hello world go")

	words := strings.Split(s, " ")
	count := len(words)
	//fmt.Println(count)
	return count
}

func jsonMe() {
	type User struct {
		Name string `json:"Name"`
		Age int `json:"Age"`
	}


}
func main() {
	//mapFunc() // name if ages > 18
	//fmt.Println(WordCount("Hello world Go")) // 3


}