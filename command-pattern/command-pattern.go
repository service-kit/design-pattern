package command_pattern

import "fmt"

type Order interface {
	Execute()
}

type Stock struct {
	Name     string
	Quantity int
}

func (s *Stock) Buy() {
	fmt.Printf("Stock [ Name: %v Quantity: %v ] Bought\n", s.Name, s.Quantity)
}

func (s *Stock) Sell() {
	fmt.Printf("Stock [ Name: %v Quantity: %v ] Sold\n", s.Name, s.Quantity)
}

type BuyStock struct {
	stock Stock
}

func (b *BuyStock) Execute() {
	b.stock.Buy()
}

type SellStock struct {
	stock Stock
}

func (b *SellStock) Execute() {
	b.stock.Sell()
}

type Broker struct {
	orderList []Order
}

func (b *Broker) TakeOrder(order Order) {
	b.orderList = append(b.orderList, order)
}

func (b *Broker) PlaceOrders() {
	for _, order := range b.orderList {
		order.Execute()
	}
	b.orderList = make([]Order, 0)
}
