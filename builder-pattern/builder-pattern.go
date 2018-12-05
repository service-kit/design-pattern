package builder_pattern

import "fmt"

type Packing interface {
	Pack() string
}

type Item interface {
	Name() string
	Packing() Packing
	Price() float32
}

type Wrapper struct {
}

func (w *Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct {
}

func (b *Bottle) Pack() string {
	return "Bottle"
}

type Burger struct {
}

func (b *Burger) Packing() Packing {
	return &Wrapper{}
}

func (b *Burger) Name() string {
	return ""
}

func (b *Burger) Price() float32 {
	return 0
}

type VegBurger struct {
	Burger
}

func (v *VegBurger) Name() string {
	return "VegBurger"
}

func (v *VegBurger) Price() float32 {
	return 25.0
}

type ChickenBurger struct {
	Burger
}

func (v *ChickenBurger) Name() string {
	return "ChickenBurger"
}

func (v *ChickenBurger) Price() float32 {
	return 50.0
}

type ColdDrink struct {
}

func (c *ColdDrink) Name() string {
	return ""
}

func (c *ColdDrink) Packing() Packing {
	return &Bottle{}
}

func (c *ColdDrink) Price() float32 {
	return 0
}

type Cock struct {
	ColdDrink
}

func (c *Cock) Price() float32 {
	return 30.0
}

func (c *Cock) Name() string {
	return "Cock"
}

type Pepsi struct {
	ColdDrink
}

func (c *Pepsi) Price() float32 {
	return 35.0
}

func (c *Pepsi) Name() string {
	return "Pepsi"
}

type Meal struct {
	items []Item
}

func (m *Meal) AddItem(i Item) {
	m.items = append(m.items, i)
}

func (m *Meal) ShowItems() {
	for _, item := range m.items {
		fmt.Print("Item ", item.Name())
		fmt.Print("\tPacking ", item.Packing().Pack())
		fmt.Println("\tPrice ", item.Price())
	}
}

func (m *Meal) GetCost() float32 {
	var cost float32 = 0
	for _, item := range m.items {
		cost += item.Price()
	}
	return cost
}

type MealBuilder struct {
}

func (m *MealBuilder) PrepareVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(&VegBurger{})
	meal.AddItem(&Cock{})
	return meal
}

func (m *MealBuilder) PrepareNonVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(&ChickenBurger{})
	meal.AddItem(&Pepsi{})
	return meal
}
