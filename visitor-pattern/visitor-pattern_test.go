package visitor_pattern

import "testing"

func Test_VisitorPattern(t *testing.T) {
	c := new(Computer)
	c.parts = []ComputerPart{new(Keyboard), new(Mouse), new(Monitor)}
	c.Accept(new(ComputerPartDisplayVisitor))
}
