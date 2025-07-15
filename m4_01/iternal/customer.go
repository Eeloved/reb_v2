package iternal

import "errors"

const DEFAULT_DISCOUNT = 500

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func (c *Customer) WrOffDebt() error {
	if c.Debt >= c.Balance {
		return errors.New("Not possible write off")
	}

	c.Balance -= c.Debt
	c.Debt = 0

	return nil
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}
