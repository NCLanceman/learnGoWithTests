package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type RightTriangle struct {
	Base, Height float64
}

func (r Rectangle) Perimeter() (result float64) {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() (result float64) {
	return 2 * math.Pi * c.Radius
}

func (t RightTriangle) Perimeter() (result float64) {
	return t.Base + t.Height + math.Sqrt((t.Base*t.Base)+(t.Height*t.Height))
}

func (r Rectangle) Area() (result float64) {
	return r.Width * r.Height
}

func (c Circle) Area() (result float64) {
	return math.Pi * c.Radius * c.Radius
}

func (t RightTriangle) Area() (result float64) {
	return 0.5 * (t.Base * t.Height)
}
