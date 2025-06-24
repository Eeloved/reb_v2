package main

import "fmt"

func main() {
	err := someFunc()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after someFunc")
}

func someFunc() (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "regular error":
				err = fmt.Errorf("application error")
			default:
				panic("critical")
			}
		}
	}()

	panic("regular error")
}

/*
	defer fmt.Println("main: defer")

	someFunc()
	fmt.Println("main: afwter someFunc call")
}

func someFunc() {
	defer fmt.Println("someFunc: defer")

	panic("something get wrong")
}
*/
