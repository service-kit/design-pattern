package flyweight_pattern

import (
	"math/rand"
	"testing"
)

func Test_FlyweightPattern(t *testing.T) {
	sf := new(ShapeFactory)
	sf.Init()
	for i := 0; i < 100; i++ {
		c := sf.GetCircle(RandColor()).(*Circle)
		c.Init(rand.Int()%100, rand.Int()%100, rand.Int()%100)
		c.Draw()
	}
}
