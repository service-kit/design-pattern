package memento_pattern

import (
	"fmt"
	"testing"
)

func Test_MementoPattern(t *testing.T) {
	c := new(CareTaker)
	o := &Originator{}
	o.SetState("state #1")
	c.Add(o.SaveStateToMemento())
	o.SetState("state #2")
	c.Add(o.SaveStateToMemento())
	o.SetState("state #3")
	c.Add(o.SaveStateToMemento())
	o.SetState("state #4")
	c.Add(o.SaveStateToMemento())
	o.SetState("state #5")
	c.Add(o.SaveStateToMemento())
	o.GetStateFromMemento(c.Get(0))
	fmt.Println(o.GetState())
	o.GetStateFromMemento(c.Get(2))
	fmt.Println(o.GetState())
	o.GetStateFromMemento(c.Get(4))
	fmt.Println(o.GetState())
}
