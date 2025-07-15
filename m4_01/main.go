package main

import (
	"fmt"
	"m4_01/iternal"
)

//const DEFAULT_DISCOUNT = 500

func main() {
	cust := iternal.NewCustomer("Dima", 23, 10000, 1000, true)

	cust.WrOffDebt()
	fmt.Printf("%+v\n", cust)
}

/*
	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not acailable")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	//discount, _ := cust.CalcDiscount()

	fmt.Printf("%+v\n", cust)
}

cust.Name = "Dima"
fmt.Printf(cust.Name)

//fmt.Printf("%+v\n", cust)*/
