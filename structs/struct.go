package structs

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct{
	Radius float64
}

func (r Rectangle) Perimeter() (result float64){
	return 2*(r.Width + r.Height) 
}

func (r Rectangle) Area() (result float64){
	return r.Width * r.Height
}

func (c Circle) Area() (result float64){
	return math.Pi * c.Radius * c.Radius 
}
