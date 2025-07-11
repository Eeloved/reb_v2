package main

import (
	"fmt"
	newcolor "myapp/packages/color"

	//"myapp/packages/utils" // имя модуля + имя папки
	. "myapp/packages/wordz"

	"github.com/fatih/color"
)

func main() {
	/*sum := utils.Add(3, 4)
	fmt.Println("Сумма:", sum)
	utils.PrintHello()*/
	newcolor.Greet()
	color.Cyan("Hello, World!")

	fmt.Println(Hello)
	fmt.Println(Random())
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(wordz.Hello)
			fmt.Println(wordz.Random())
		}*/
}

/*	func main() {
	fmt.Println("Hello, World!")
	//color.Red("Hello, World!")

	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())
	}
}*/
