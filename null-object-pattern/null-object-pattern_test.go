package null_object_pattern

import (
	"fmt"
	"testing"
)

func Test_NullObjectPattern(t *testing.T) {
	cf := new(CustomerFactory)
	cf.Init()
	cf.AddCustomer(&RealCustomer{Name: "kevin"})
	cf.AddCustomer(&RealCustomer{Name: "lucy"})
	fmt.Println(cf.GetCustomer("kevin").GetName())
	fmt.Println(cf.GetCustomer("lucy").GetName())
	fmt.Println(cf.GetCustomer("jack").GetName())
}
