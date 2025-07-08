package main

import (
	"fmt"
	"m4_01/iternal"
)


func main() {
	cust := iternal.Customer{
		Age: 23,
		Balance: 100000,
		Debt: 1000,
		Name: "Vasilii",
	}
	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not acailable")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0{
			return 0, nil
		}
		return result, nil
		}
		discount, _ := cust.CalcDiscount()

		fmt.Printf("%d", discount)
	}

	/*cust.Name = "Dima"
	fmt.Printf(cust.Name)

	//fmt.Printf("%+v\n", cust)*/
