package bridge_pattern

import "fmt"

type DrawAPI interface {
	DrawCircle(redius, x, y int)
}

type RedCircle struct {
}

func (r *RedCircle) DrawCircle(redius, x, y int) {
	fmt.Printf("Draw Red Circle redius:%v x:%v y:%v\n", redius, x, y)
}

type GreenCircle struct {
}

func (r *GreenCircle) DrawCircle(redius, x, y int) {
	fmt.Printf("Draw Green Circle redius:%v x:%v y:%v\n", redius, x, y)
}

type Shape struct {
	drawApi DrawAPI
}

func (s *Shape) Init(api DrawAPI) {
	s.drawApi = api
}

func (s *Shape) Draw(redius, x, y int) {
}

type Circle struct {
	Shape
}

func (c *Circle) Init(api DrawAPI) {
	c.Shape.Init(api)
}

func (c *Circle) Draw(redius, x, y int) {
	c.drawApi.DrawCircle(redius, x, y)
}
