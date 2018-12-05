package facade_pattern

import "testing"

func Test_FacadePattern(t *testing.T) {
	m := new(ShapeMaker)
	m.Init()
	m.DrawCircle()
	m.DrawRectangle()
	m.DrawSquare()
}
