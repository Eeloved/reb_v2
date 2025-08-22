package shapes


type Rectangle struct {
	Width  float64
	Height float64
}

type Areaer interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64{
	return 3.14159 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}
