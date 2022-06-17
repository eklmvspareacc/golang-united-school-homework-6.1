package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (copy Triangle) CalcPerimeter() float64 {
	return 3 * copy.Side
}

func (copy Triangle) CalcArea() float64 {
	return (math.Sqrt(3) / 4) * copy.Side * copy.Side
}
