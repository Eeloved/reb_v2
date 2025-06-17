package main

import "fmt"

func main() {
	slice := []string{"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}
	//fmt.Println(len(slice), cap(slice))
	budny := []string{}
	vihi := []string{}
	toogether := []string{}
	for _, day := range slice {
		if day == "суббота" || day == "воскресенье" {
			vihi = append(vihi, day)
		} else {
			budny = append(budny, day)
		}
	}
	slice = budny

	fmt.Println(cap(slice), len(slice), budny, cap(vihi), len(vihi), vihi)

	toogether = append(slice, vihi...)
	fmt.Println(toogether)
}
