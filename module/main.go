/*
package main

import (

	"fmt"
	"module\src\gopackages\wordz"

	"github.com/fatih/color"

)

	func main() {
		fmt.Println("Hello, World!")
		color.Red("Hello, World!")

		for i := 0; i < 5; i++ {
			fmt.Println(wordz.Hello)
			fmt.Println(wordz.Random())
		}
	}
*/
package main

import (
	"fmt"
	"module/utils" // имя модуля + имя папки
)

func main() {
	sum := utils.Add(3, 4)
	fmt.Println("Сумма:", sum)
	utils.PrintHello()
}
