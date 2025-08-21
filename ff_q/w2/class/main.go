package main

import (
	"fmt"
	"myapp/utils"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s, and I'm %d years old.\n", p.Name, p.Age)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

type Speaker interface {
	Speak() string
}

func (p Person) Speak() string {
	return fmt.Sprintf("My name %s", p.Name)
}
func main() {
	p1 := Person{Name: "Alice", Age: 30}
	//p2 := Person{"Bob", 25} //not recommended

	fmt.Println(p1.Name)
	p1.Age = 31
	p1.Greet()

	p1.HaveBirthday()
	p1.Greet()

	fmt.Println(utils.IsEven(4))

	fmt.Println(p1.Speak())
}
