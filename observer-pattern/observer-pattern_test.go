package observer_pattern

import "testing"

func Test_ObserverPattern(t *testing.T) {
	s := new(Subject)
	b := new(BinaryObserver)
	b.Init(s)
	c := new(OctalObserver)
	c.Init(s)
	s.SetState("1")
	s.SetState("2")
}
