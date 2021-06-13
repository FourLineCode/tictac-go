package game

import (
	"github.com/FourLineCode/tictac-go/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	cfg     config.Config = config.GetConfig()
	xOffset float32       = float32(cfg.WindowWidth-int32(boardSize)) / 2
	yOffset float32       = float32(cfg.WindowHeight-int32(boardSize)) / 2
)

func getVector2(x, y float32) rl.Vector2 {
	xpos := x + xOffset
	ypos := y + yOffset

	return rl.Vector2{X: xpos, Y: ypos}
}
