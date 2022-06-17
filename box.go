package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorOutOfCapacity   = errors.New("maximum capacity reached")
	errorIndexOutOfRange = errors.New("index out of range")
	errorNoCirclesFound  = errors.New("no circles found")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapes:         make([]Shape, 0, shapesCapacity),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf("%w", errorOutOfCapacity)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if !b.isIndexInRange(i) {
		return nil, fmt.Errorf("%w", errorIndexOutOfRange)
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	old, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	b.shapes[i] = shape
	return old, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := 0.0
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum := 0.0
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	result := make([]Shape, 0, b.shapesCapacity)
	foundCircle := false
	for _, shape := range b.shapes {
		_, ok := shape.(Circle)
		if ok {
			foundCircle = true
			continue
		}
		result = append(result, shape)
	}
	if foundCircle {
		b.shapes = result
		return nil
	}
	return fmt.Errorf("%w", errorNoCirclesFound)
}

func (b *box) isIndexInRange(i int) bool {
	return i < len(b.shapes)
}
