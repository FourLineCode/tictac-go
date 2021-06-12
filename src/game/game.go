package game

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
}

func New() *Game {
	return &Game{}
}

func (g *Game) Draw() {
	g.drawInfo()
}

func (g *Game) drawInfo() {
	drawFPS()
	drawWindowSize()
	drawElapsedTime()
}

func drawFPS() {
	currentFPS := rl.GetFPS()

	fpsText := fmt.Sprintf("FPS: %v\n", currentFPS)

	rl.DrawText(fpsText, 5, 5, 12, rl.Green)
}

func drawWindowSize() {
	currentWidth := rl.GetScreenWidth()
	currentHeight := rl.GetScreenHeight()

	dimensionText := fmt.Sprintf("Window: %v x %v\n", currentWidth, currentHeight)

	rl.DrawText(dimensionText, 5, 17, 12, rl.White)
}

func drawElapsedTime() {
	currentTime := rl.GetTime()
	duration, _ := time.ParseDuration(fmt.Sprintf("%fs", currentTime))

	mins := int64(duration.Minutes()) % 60
	secs := int64(duration.Seconds()) % 60
	mils := duration.Milliseconds() % 100

	timeText := fmt.Sprintf("El. Time: %02d : %02d : %02d\n", mins, secs, mils)

	rl.DrawText(timeText, 5, 29, 12, rl.White)
}
