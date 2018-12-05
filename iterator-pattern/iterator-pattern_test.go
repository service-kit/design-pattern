package iterator_pattern

import (
	"fmt"
	"testing"
)

func Test_IteratorPattern(t *testing.T) {
	list := NewList()
	list.Pushback("1")
	list.Pushback("2")
	list.Pushback("3")
	list.Pushback("4")
	list.Pushback("5")
	list.Pushback("6")
	list.Pushback("7")
	list.Pushback("8")
	list.Popback()
	list.Popback()
	for iter := list.First(); iter != list.End(); iter = iter.Next() {
		fmt.Println(iter.Value())
	}
}
