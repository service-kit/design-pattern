package composite_pattern

import (
	"fmt"
	"testing"
)

func Test_CompositePattern(t *testing.T) {
	emp := new(Employee)
	emp.Name = "CEO"
	te := Employee{Name: "team leader"}
	te.Add(Employee{Name: "lilei"})
	te.Add(Employee{Name: "jack"})
	te.Add(Employee{Name: "lily"})
	te.Add(Employee{Name: "lucy"})
	te.Add(Employee{Name: "john"})
	emp.Add(te)
	fmt.Println(emp)
}
