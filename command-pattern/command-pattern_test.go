package command_pattern

import "testing"

func Test_CommandPattern(t *testing.T) {
	bs := &BuyStock{stock: Stock{Name: "dog", Quantity: 3000}}
	ss := &SellStock{stock: Stock{Name: "cat", Quantity: 1000}}
	b := &Broker{}
	b.TakeOrder(bs)
	b.TakeOrder(ss)
	b.PlaceOrders()
}
