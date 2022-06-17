package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (copy Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * copy.Radius
}

func (copy Circle) CalcArea() float64 {
	return math.Pi * copy.Radius * copy.Radius
}
