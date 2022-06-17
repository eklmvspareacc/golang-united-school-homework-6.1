package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (copy Rectangle) CalcPerimeter() float64 {
	return 2 * (copy.Height + copy.Weight)
}

func (copy Rectangle) CalcArea() float64 {
	return copy.Height * copy.Weight
}
