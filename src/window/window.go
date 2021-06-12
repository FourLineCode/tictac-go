package window

import (
	"github.com/FourLineCode/tictac-go/internal/config"
	"github.com/FourLineCode/tictac-go/src/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type window struct {
	config config.Config
	game   *game.Game
}

func New(cfg config.Config) *window {
	return &window{
		config: cfg,
		game:   game.New(),
	}
}

func (w *window) Run() {
	rl.InitWindow(w.config.WindowWidth, w.config.WindowHeight, w.config.WindowTitle)

	rl.SetTargetFPS(w.config.WindowFPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(w.config.WindowBackgroundColor)

		w.drawInfo()

		w.game.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
