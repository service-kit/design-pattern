package builder_pattern

import (
	"fmt"
	"testing"
)

func Test_BuilderPattern(t *testing.T) {
	mb := new(MealBuilder)
	vm := mb.PrepareVegMeal()
	nm := mb.PrepareNonVegMeal()
	vm.ShowItems()
	fmt.Println("total cost:", vm.GetCost())
	nm.ShowItems()
	fmt.Println("total cost:", nm.GetCost())
}
