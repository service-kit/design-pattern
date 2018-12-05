package facade_pattern

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Rectangle::Draw")
}

type Square struct {
}

func (r *Square) Draw() {
	fmt.Println("Square::Draw")
}

type Circle struct {
}

func (r *Circle) Draw() {
	fmt.Println("Circle::Draw")
}

type ShapeMaker struct {
	circle    Shape
	rectangle Shape
	square    Shape
}

func (s *ShapeMaker) Init() {
	s.circle = new(Circle)
	s.rectangle = new(Rectangle)
	s.square = new(Square)
}

func (s *ShapeMaker) DrawCircle() {
	s.circle.Draw()
}

func (s *ShapeMaker) DrawRectangle() {
	s.rectangle.Draw()
}

func (s *ShapeMaker) DrawSquare() {
	s.square.Draw()
}
