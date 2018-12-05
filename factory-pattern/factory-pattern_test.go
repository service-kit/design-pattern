package factory_pattern

import "testing"

func Test_FactoryPattern(t *testing.T) {
	f := new(ShapeFactory)
	r := f.GetShape(RECTANGLE)
	r.Draw()
	s := f.GetShape(SQUARE)
	s.Draw()
	c := f.GetShape(CIRCLE)
	c.Draw()
	t.Log("test factory pattern pass")
}

func Test_AbstractFactoryPattern(t *testing.T) {
	fp := new(FactoryProducer)
	cf := fp.GetFactory(COLOR)
	c := cf.GetColor(RED)
	c.Fill()
	sf := fp.GetFactory(SHAPE)
	s := sf.GetShape(CIRCLE)
	s.Draw()
	t.Log("test factory producer pass")
}
