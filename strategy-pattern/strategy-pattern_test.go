package strategy_pattern

import (
	"fmt"
	"testing"
)

func Test_StrategyPattern(t *testing.T) {
	c := new(Context)
	c.SetStrategy(&OperationAdd{})
	fmt.Println(c.ExecuteStrategy(5, 2))
	c.SetStrategy(&OperationSubstract{})
	fmt.Println(c.ExecuteStrategy(5, 2))
}
