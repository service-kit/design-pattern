package state_pattern

import "fmt"

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) GetState() State {
	return c.state
}

type State interface {
	DoAction(*Context)
}

type StartState struct {
}

func (s *StartState) DoAction(context *Context) {
	fmt.Println("start state")
	context.SetState(s)
}

type StopState struct {
}

func (s *StopState) DoAction(context *Context) {
	fmt.Println("stop state")
	context.SetState(s)
}
