package decorator_pattern

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Rectangle::Draw")
}

type Circle struct {
}

func (r *Circle) Draw() {
	fmt.Println("Circle::Draw")
}

type ShapeDecorator struct {
	decoratedShape Shape
}

func (s *ShapeDecorator) Draw() {
	s.decoratedShape.Draw()
}

type RedShapeDecorator struct {
	ShapeDecorator
}

func (r *RedShapeDecorator) Draw() {
	r.decoratedShape.Draw()
	r.SetRedBorder(r.decoratedShape)
}

func (r *RedShapeDecorator) SetRedBorder(s Shape) {
	fmt.Println("Border Color: Red")
}
