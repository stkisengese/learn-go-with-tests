package structsmethods

import "math"

type Rectangle struct {
	Length float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Length + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
