package main

import "fmt"

func main() {
	var AmericanVelocity, EuropeanVelocity float64

	EuropeanVelocity = 102.4 * 60 * 60 / 1000 // m/s -> km/h
	AmericanVelocity = 130 * 60 * 60 / 1609   // m/s -> mph
	fmt.Println(EuropeanVelocity, AmericanVelocity)
}
