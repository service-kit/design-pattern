package strategy_pattern

type Strategy interface {
	Operation(int, int) int
}

type OperationAdd struct {
}

func (a *OperationAdd) Operation(num1, num2 int) int {
	return num1 + num2
}

type OperationSubstract struct {
}

func (a *OperationSubstract) Operation(num1, num2 int) int {
	return num1 - num2
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.Operation(num1, num2)
}
