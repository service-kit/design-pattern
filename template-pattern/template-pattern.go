package template_pattern

import "fmt"

type IGame interface {
	Init()
	StartPlay()
	FinishPlay()
}

type Game struct {
	IGame
}

func (g *Game) Play() {
	g.Init()
	g.StartPlay()
	g.FinishPlay()
}

type Cricket struct {
	Game
}

func (c *Cricket) Init() {
	fmt.Println("init cricket")
}

func (c *Cricket) StartPlay() {
	fmt.Println("start play cricket")
}

func (c *Cricket) FinishPlay() {
	fmt.Println("finish play cricket")
}

func NewCricket() Game {
	c := new(Cricket)
	c.IGame = c
	return c.Game
}

type Football struct {
	Game
}

func (c *Football) Init() {
	fmt.Println("init football")
}

func (c *Football) StartPlay() {
	fmt.Println("start play football")
}

func (c *Football) FinishPlay() {
	fmt.Println("finish play football")
}

func NewFootball() Game {
	c := new(Football)
	c.IGame = c
	return c.Game
}
