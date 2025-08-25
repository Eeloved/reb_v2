package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func mapFunc() {

	ages := map[string]int{
		"Alice":   25,
		"Bob":     17,
		"Eve":     30,
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
		Age  int    `json:"Age"`
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

func saveJson() {

	type User struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}
	users := []User{
		{Name: "Alice", Age: 25},
		{Name: "Jhon", Age: 33},
		{Name: "Alex", Age: 12},
	}

	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("users.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readJson() {

	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	var loadedUsers []User
	err = json.Unmarshal(data, &loadedUsers)
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	for _, users := range loadedUsers {
		names = append(names, users.Name)
		fmt.Println("Name:", users.Name)
	}
	fmt.Println("Names:", names)
}

func ParseAge(ageStr string) (int, error) {
	return strconv.Atoi(ageStr)

}

func main() {
	//mapFunc() // name if ages > 18
	//fmt.Println(WordCount("Hello world Go")) // 3
	//jsonMe()
	//saveJson()
	//readJson()

}

/*    data, err := json.Marshal(users)
if err != nil {
    log.Fatal(err)
} */
