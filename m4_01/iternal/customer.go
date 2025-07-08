package iternal

type Customer struct {
	Name   string
	Age    int
	Balance int
	Debt int
	Discount bool
	CalcDiscount func() (int, error)
}


