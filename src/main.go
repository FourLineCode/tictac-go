package main

import (
	"github.com/FourLineCode/tictac-go/internal/config"
	"github.com/FourLineCode/tictac-go/src/window"
)

const (
	WIDTH  = 800
	HEIGHT = 600
)

func main() {
	cfg := config.GetConfig()

	win := window.New(cfg)

	win.Run()
}
