package mvc_pattern

import "testing"

func Test_MVCPattern(t *testing.T) {
	s := &Student{"kevin", 18}
	v := &StudentView{s}
	c := &StudentController{s}
	v.Show()
	c.AddAge()
	v.Show()
}
