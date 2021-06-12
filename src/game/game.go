package game

import "fmt"

type Game struct {
}

func New() *Game {
	return &Game{}
}

func (g *Game) Draw() {
	fmt.Println("Drawing Game Objects")
}
