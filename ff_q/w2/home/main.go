package main

import (
	"fmt"
	"myapp/shapes"
)

func printArea(a shapes.Areaer) {
	fmt.Println("ploshad", a.Area())
}

func main() {
	r := shapes.Rectangle{Width: 10, Height: 5}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	r.Scale(2)
	fmt.Println(r.Area())

	printArea(r)

	c := shapes.Circle{Radius: 5}
	fmt.Println(c.Area())
	printArea(c)
}
