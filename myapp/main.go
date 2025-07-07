package main

import (
	"fmt"
	"myapp/utils" // имя модуля + имя папки
	"myapp/wordz"

	"github.com/fatih/color"
)

func main() {
	sum := utils.Add(3, 4)
	fmt.Println("Сумма:", sum)
	utils.PrintHello()

	color.Cyan("Hello, World!")

	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())
	}
}

/*	func main() {
	fmt.Println("Hello, World!")
	//color.Red("Hello, World!")

	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())
	}
}*/
