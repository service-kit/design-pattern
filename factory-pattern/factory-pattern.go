package factory_pattern

import "fmt"

const (
	//Shape Type
	CIRCLE    = "CIRCLE"
	RECTANGLE = "RECTANGLE"
	SQUARE    = "SQUARE"
	//Color Type
	RED   = "RED"
	GREEN = "GREEN"
	BLUE  = "BLUE"
	//Factory Type
	SHAPE = "SHAPE"
	COLOR = "COLOR"
)

type Shape interface {
	Draw()
}

type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Circle::Draw()")
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Rectangle::Draw()")
}

type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("Square::Draw()")
}

type ShapeFactory struct {
}

func (f *ShapeFactory) GetShape(s string) Shape {
	if CIRCLE == s {
		return &Circle{}
	} else if RECTANGLE == s {
		return &Rectangle{}
	} else if SQUARE == s {
		return &Square{}
	} else {
		return nil
	}
}

func (f *ShapeFactory) GetColor(c string) Color {
	return nil
}

type Color interface {
	Fill()
}

type Red struct {
}

func (r *Red) Fill() {
	fmt.Println("Red::Fill")
}

type Green struct {
}

func (g *Green) Fill() {
	fmt.Println("Green::Fill")
}

type Blue struct {
}

func (b *Blue) Fill() {
	fmt.Println("Blue::Fill")
}

type ColorFactory struct {
}

func (f *ColorFactory) GetColor(c string) Color {
	if RED == c {
		return &Red{}
	} else if GREEN == c {
		return &Green{}
	} else if BLUE == c {
		return &Blue{}
	} else {
		return nil
	}
}

func (f *ColorFactory) GetShape(s string) Shape {
	return nil
}
