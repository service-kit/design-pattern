package decorator_pattern

import (
	"testing"
)

func Test_DecoratorPattern(t *testing.T) {
	c := &Circle{}
	rc := &RedShapeDecorator{ShapeDecorator{decoratedShape: &Circle{}}}
	rr := &RedShapeDecorator{ShapeDecorator{decoratedShape: &Rectangle{}}}
	c.Draw()
	rc.Draw()
	rr.Draw()
}
