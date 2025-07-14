package main

import(
	"fmt"
	"time"
)

type Buyer interface {
	buy()
}

type Seller interface{
	sell()
}
type Dealer interface{
	Buyer
	Seller
}

type Checker interface{
	CheckDate() bool
}
type Card struct {
	Balance int
	ExpitedDate string
	CVV int
	Num int64
	Owner string
}

type CreditCard struct {
	*Card
	Limit int
}

func (c *Card) CheckDate() bool {
	ex, _ := time.Parse("2006.01.02", c.ExpitedDate)
	return time.Now().Before(ex)
}

func main() {
		c := &Card{
			Balance: 10000,
			ExpitedDate: "01.02.2027",
			CVV: 132,
			Num: 4234536475474656,
			Owner: "Petr Velikii",
		}
		cc := &CreditCard{
			c,
			100000,
		}
		fmt.Printf("%t\n", cc.CheckDate())
	}/*
	fmt.Printf("%+v\n", cc)

	fmt.Printf("Limit: %d, Owner: %s\n", cc.Limit, cc.Owner)

	cc.Owner = "Oleg"
	cc.Limit = 0

	fmt.Printf("Limit: %d, Owner: %s\n", cc.Limit, cc.Owner)
	*/
