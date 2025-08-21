package main

import "fmt"

func greet(name string) {
	fmt.Printf("Привет, %s!\n", name)
}

func isAdult(age int) bool {
	return age >= 18

}
func main() {
	var name string
	var age int
	fmt.Print("Как тебя зовут? ")
	fmt.Scan(&name)
	fmt.Print("Сколько тебе лет? ")
	fmt.Scan(&age)

	greet(name)
	if isAdult(age) {
		fmt.Println("Ты взрослый")
	} else {
		fmt.Println("Ты не возрослый")
	}
}

/*
	var name string
	var age int

	fmt.Print("Как тебя зовут? ")
	fmt.Scan(&name)
	fmt.Print("Сколько тебе лет? ")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Printf("Извини, %s, ты слишком молод!\n", name)
	} else if age < 60 {
		fmt.Printf("Добро пожаловать, %s!\n", name)
	} else {
		fmt.Printf("%s, ты имеешь право на скидку!\n", name)
	}

	fmt.Println("Программа завершена")
}
*/
