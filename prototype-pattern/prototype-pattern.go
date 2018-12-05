package prototype_pattern

import (
	"fmt"
	"github.com/service-kit/design-pattern/util"
)

type Cloneable interface {
	Clone() interface{}
}

type IShape interface {
	Draw()
	Clone() interface{}
	GetType() string
	GetID() string
}

type Shape struct {
	ID   string
	Type string
}

func (s *Shape) Draw() {
}

func (s *Shape) Clone() interface{} {
	return util.Copy(s)
}

func (s *Shape) GetType() string {
	return s.Type
}

func (s *Shape) GetID() string {
	return s.ID
}

func (s *Shape) SetType(t string) {
	s.Type = t
}

func (s *Shape) SetID(id string) {
	s.ID = id
}

type Rectangle struct {
	Shape
}

func (r *Rectangle) Draw() {
	fmt.Println("Rectangle::Draw")
}

type Square struct {
	Shape
}

func (r *Square) Draw() {
	fmt.Println("Square::Draw")
}

type Circle struct {
	Shape
}

func (r *Circle) Draw() {
	fmt.Println("Circle::Draw")
}

type ShapeCache struct {
	ShapeMap map[string]IShape
}

func (s *ShapeCache) LoadCache() {
	s.ShapeMap = make(map[string]IShape)
	c := &Circle{Shape{Type: "Circle"}}
	c.SetID("1")
	s.ShapeMap[c.GetID()] = c
	r := &Rectangle{Shape{Type: "Rectangle"}}
	r.SetID("2")
	s.ShapeMap[r.GetID()] = r
	sq := &Square{Shape{Type: "Square"}}
	sq.SetID("3")
	s.ShapeMap[sq.GetID()] = sq
}

func (s *ShapeCache) GetShape(id string) IShape {
	return s.ShapeMap[id].Clone().(IShape)
}
