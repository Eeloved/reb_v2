package main

import (
	"fmt"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Aera() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	fmt.Println(r.Aera())
	fmt.Println(r.Perimeter())
}
