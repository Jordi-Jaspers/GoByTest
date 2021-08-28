package interfaces

import "math"

// This is quite different to interfaces in most other programming languages.
// Normally you have to write code to say My type Foo implements interface Bar.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ========================== RECTANGLE ========================== //
type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

// ========================== Circle ========================== //
type Circle struct {
	Radius float64
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}
