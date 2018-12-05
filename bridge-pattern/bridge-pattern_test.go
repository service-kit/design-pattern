package bridge_pattern

import "testing"

func Test_BridgePattern(t *testing.T) {
	c1 := Circle{Shape{drawApi: &RedCircle{}}}
	c2 := Circle{Shape{drawApi: &GreenCircle{}}}
	c1.Draw(10, 10, 10)
	c2.Draw(10, 10, 10)
}
