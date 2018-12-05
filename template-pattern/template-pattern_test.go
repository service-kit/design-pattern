package template_pattern

import "testing"

func Test_TemplatePattern(t *testing.T) {
	c := NewCricket()
	c.Play()
	f := NewFootball()
	f.Play()
}
