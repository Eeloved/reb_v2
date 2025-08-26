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

	for _, u := range loadedUsers {
        log.Println("Взрослый:", u.Name)
		
	}
}
func ParseAge(ageStr string) (int, error) {
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        return 0, err
    }
    return age, nil
}

func readFileCount(){
    // Чтение файла
    content, err := os.ReadFile("text.txt")
    if err != nil {
        fmt.Println("Ошибка чтения файла:", err)
        return
    }

    // Преобразование содержимого в строку и разделение на слова
    text := string(content)
    words := strings.Fields(text)

    // Подсчет слов
    wordCount := len(words)

    // Вывод результата
    fmt.Printf("Количество слов в файле: %d\n", wordCount)
}

func configJson(){

	type Config struct {
		Port    int    `json:"port"`
		Host    string `json:"host"`
		Debug   bool   `json:"debug"`
	}


    // Create config.json file
    config := Config{
        Port:  8080,
        Host:  "localhost",
        Debug: true,
    }

    // Write config to file
    configData, err := json.MarshalIndent(config, "", "    ")
    if err != nil {
        fmt.Printf("Error marshaling config: %v\n", err)
        return
    }

    err = os.WriteFile("config.json", configData, 0644)
    if err != nil {
        fmt.Printf("Error writing config file: %v\n", err)
        return
    }

    // Read config from file
    fileData, err := os.ReadFile("config.json")
    if err != nil {
        fmt.Printf("Error reading config file: %v\n", err)
        return
    }

    // Unmarshal JSON into Config struct
    var loadedConfig Config
    err = json.Unmarshal(fileData, &loadedConfig)
    if err != nil {
        fmt.Printf("Error unmarshaling config: %v\n", err)
        return
    }

    // Print loaded config values
    fmt.Printf("Port: %d\n", loadedConfig.Port)
    fmt.Printf("Host: %s\n", loadedConfig.Host)
    fmt.Printf("Debug: %v\n", loadedConfig.Debug)
}

func main() {
	//mapFunc() // name if ages > 18
	//fmt.Println(WordCount("Hello world Go")) // 3
	//jsonMe()
	//saveJson()
	//readJson()


/*	fmt.Println(ParseAge("adb"))
	age, err := ParseAge("25")
    if err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Println("Возраст:", age)
    } */

	//readFileCount()
	configJson()
}

/*    data, err := json.Marshal(users)
if err != nil {
    log.Fatal(err)
} */
