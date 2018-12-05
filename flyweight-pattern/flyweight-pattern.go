package flyweight_pattern

import (
	"fmt"
	"math/rand"
)

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
	WHITE = "white"
	BLACK = "black"
)

var COLORS = []string{RED, GREEN, BLUE, WHITE, BLACK}

type Shape interface {
	Draw()
}

type Circle struct {
	color  string
	x      int
	y      int
	radius int
}

func (c *Circle) Init(x, y, radius int) {
	c.x, c.y, c.radius = x, y, radius
}

func (c *Circle) Draw() {
	fmt.Printf("Circle::Draw color:%v x:%v y:%v radius:%v\n", c.color, c.x, c.y, c.radius)
}

func RandColor() string {
	return COLORS[int(rand.Uint32())%len(COLORS)]
}

type ShapeFactory struct {
	shapeMap map[string]Shape
}

func (s *ShapeFactory) Init() {
	s.shapeMap = make(map[string]Shape)
}

func (s *ShapeFactory) GetCircle(color string) Shape {
	if nil == s.shapeMap[color] {
		c := new(Circle)
		c.color = color
		s.shapeMap[color] = c
	}
	return s.shapeMap[color]
}
