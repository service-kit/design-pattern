package state_pattern

import "testing"

func Test_StatePattern(t *testing.T) {
	context := new(Context)
	startState := new(StartState)
	startState.DoAction(context)
	stopState := new(StopState)
	stopState.DoAction(context)
}
