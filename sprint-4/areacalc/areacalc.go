package areacalc

import "strings"

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	a    float64
	b    float64
	name string
}

func NewRectangle(a, b float64, name string) *Rectangle {
	// TODO: implement me
	return &Rectangle{
		a:    a,
		b:    b,
		name: name,
	}
}

type Circle struct {
	r    float64
	name string
}

func NewCircle(r float64, name string) *Circle {
	return &Circle{
		r:    r,
		name: name,
	}
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.a * rectangle.b
}

func (circle Circle) Area() float64 {
	return pi * circle.r * circle.r
}

func (rectangle Rectangle) Type() string {
	return rectangle.name
}

func (circle Circle) Type() string {
	return circle.name
}

func AreaCalculator(figures []Shape) (string, float64) {
	var sb strings.Builder
	var sum float64

	for i, f := range figures {
		sb.WriteString(f.Type())
		if i < len(figures)-1 {
			sb.WriteString("-")
		}
		sum += f.Area()
	}

	return sb.String(), sum
}
