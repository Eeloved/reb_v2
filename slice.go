package main

import "fmt"

func main() {
	slice := []string{"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}

	budny := []string{}
	vihi := []string{}

	for _, day := range slice {
		if day == "суббота" || day == "воскресенье" {
			budny = append(budny, day)
		} else {
			vihi = append(vihi, day)
		}
	}
	slice = vihi

	fmt.Println(budny, vihi)
}
