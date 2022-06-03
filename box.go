package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf("out of range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	valid, err := indexIsValid(b.shapes, i)
	if !valid {
		return nil, err
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	valid, err := indexIsValid(b.shapes, i)
	if !valid {
		return nil, err
	}

	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	valid, err := indexIsValid(b.shapes, i)
	if !valid {
		return nil, err
	}

	removedShape := b.shapes[i]
	b.shapes[i] = shape
	return removedShape, nil
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
	atLeastOne := false
	for i, _ := range b.shapes {
		if b.shapes[i] != nil {
			atLeastOne = true
			b.shapes[i] = nil
		}
	}

	if !atLeastOne {
		return fmt.Errorf("there are no circles in the list")
	}
	return nil
}

func indexIsValid(shapes []Shape, i int) (bool, error) {
	if i >= len(shapes) || shapes[i] == nil {
		return true, fmt.Errorf("not a valid index")
	}
	return false, nil
}
