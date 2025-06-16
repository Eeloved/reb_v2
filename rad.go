package main

import (
	"fmt"
	"math"
)

func RadiusValue() {
	// Заданная длина окружности
	var length float64 = 35

	// Вычисляем радиус окружности: R = C / (2 * π)
	radiusValue := length / (2 * math.Pi)

	// Объявляем указатель на float64 и присваиваем адрес переменной radiusValue
	var R *float64 = &radiusValue

	// Вычисляем площадь круга через разыменование указателя R
	area := math.Pi * (*R) * (*R)

	// Округляем площадь до двух знаков после запятой
	area = math.Round(area*100) / 100

	// Выводим результаты
	fmt.Printf("Радиус: %.2f\n", *R)
	fmt.Printf("Площадь круга: %.2f\n", area)
}

func main() {
	RadiusValue()
}
