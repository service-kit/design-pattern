package factory_pattern

type AbstractFactory interface {
	GetColor(c string) Color
	GetShape(s string) Shape
}

type FactoryProducer struct {
}

func (fp *FactoryProducer) GetFactory(f string) AbstractFactory {
	if SHAPE == f {
		return &ShapeFactory{}
	} else if COLOR == f {
		return &ColorFactory{}
	} else {
		return nil
	}
}
