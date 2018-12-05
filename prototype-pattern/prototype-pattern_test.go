package prototype_pattern

import (
	"fmt"
	"testing"
)

func Test_PrototypePattern(t *testing.T) {
	sc := new(ShapeCache)
	sc.LoadCache()
	s1 := sc.GetShape("1")
	s2 := sc.GetShape("2")
	s3 := sc.GetShape("3")
	fmt.Println(s1.GetType())
	fmt.Println(s2.GetType())
	fmt.Println(s3.GetType())
}
