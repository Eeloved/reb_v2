package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "104"
	b := 35

	intA, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Ошибка преобразования", err)
		return
	}
	fmt.Println(intA)

	strB := strconv.Itoa(b)
	fmt.Println(strB)

}
