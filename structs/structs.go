package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Shape interface {
	Area() float64
}

func Perimiter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}