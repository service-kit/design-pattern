package visitor_pattern

import (
	"fmt"
	"reflect"
)

var (
	KEYBOARD_TYPE = reflect.ValueOf(&Keyboard{}).Type()
	MOUSE_TYPE    = reflect.ValueOf(&Mouse{}).Type()
	MONITOR_TYPE  = reflect.ValueOf(&Monitor{}).Type()
	COMPUTER_TYPE = reflect.ValueOf(&Computer{}).Type()
)

type ComputerPartVisitor interface {
	Visit(ComputerPart)
}

type ComputerPart interface {
	Accept(ComputerPartVisitor)
}

type Keyboard struct {
}

func (k *Keyboard) Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.Visit(k)
}

type Mouse struct {
}

func (m *Mouse) Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.Visit(m)
}

type Monitor struct {
}

func (m *Monitor) Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.Visit(m)
}

type Computer struct {
	parts []ComputerPart
}

func (c *Computer) Accept(computerPartVisitor ComputerPartVisitor) {
	for _, part := range c.parts {
		computerPartVisitor.Visit(part)
	}
	computerPartVisitor.Visit(c)
}

type ComputerPartDisplayVisitor struct {
}

func (c *ComputerPartDisplayVisitor) Visit(com ComputerPart) {
	out := ""
	ct := reflect.ValueOf(com).Type()
	if ct == KEYBOARD_TYPE {
		out = "keyboard"
	} else if ct == MOUSE_TYPE {
		out = "mouse"
	} else if ct == MONITOR_TYPE {
		out = "monitor"
	} else if ct == COMPUTER_TYPE {
		out = "computer"
	}
	fmt.Println("visit ", out)
}
