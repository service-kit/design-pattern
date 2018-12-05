package chain_of_responsibility_pattern

import "testing"

func Test_ChainOfResponsibilityPattern(t *testing.T) {
	cl := NewConsoleLogger()
	fl := NewFileLogger()
	el := NewErrorLogger()
	el.SetNextLogger(fl)
	fl.SetNextLogger(cl)
	el.LogMessage(INFO, "log info")
	el.LogMessage(DEBUG, "log debug")
	el.LogMessage(ERROR, "log error")
}
