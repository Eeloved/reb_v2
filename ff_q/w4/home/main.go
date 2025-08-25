package main

import (
	"fmt"
	"strings"
	"encoding/json"
	"log"
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

	users := User{Name: "Alice", Age: 25}
/*	users := []User{
		{Name: "Alice", Age: 25},
		{Name: "Jhon", Age: 33},
		{Name: "Alex", Age: 12},
	}
*/
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json:", string(data))
	//fmt.Println("Json:", data)
	//fmt.Println(users)

    var loadedUsers User
	err = json.Unmarshal(data, &loadedUsers)
    if err != nil {
        log.Fatal(err)
	}
	//fmt.Println("Struct:", loadedUsers)
	fmt.Printf("Имя: %s, Возраст: %d\n", loadedUsers.Name, loadedUsers.Age)
	
}

func main() {
	//mapFunc() // name if ages > 18
	//fmt.Println(WordCount("Hello world Go")) // 3
	jsonMe()


}


/*    data, err := json.Marshal(users)
    if err != nil {
        log.Fatal(err)
    } */