package window

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (w *window) drawInfo() {
	drawFPS()
	drawWindowSize()
	drawElapsedTime()
}

var (
	fontSize int32 = 12
	offset   int32 = 5
)

func getTextPos(line int32) int32 {
	return line*fontSize + offset
}

func drawFPS() {
	currentFPS := rl.GetFPS()

	fpsText := fmt.Sprintf("FPS: %v\n", currentFPS)

	rl.DrawText(fpsText, offset, getTextPos(0), 12, rl.Green)
}

func drawWindowSize() {
	currentWidth := rl.GetScreenWidth()
	currentHeight := rl.GetScreenHeight()

	dimensionText := fmt.Sprintf("Window: %v x %v\n", currentWidth, currentHeight)

	rl.DrawText(dimensionText, offset, getTextPos(1), 12, rl.White)
}

func drawElapsedTime() {
	currentTime := rl.GetTime()
	duration, _ := time.ParseDuration(fmt.Sprintf("%fs", currentTime))

	mins := int64(duration.Minutes()) % 60
	secs := int64(duration.Seconds()) % 60
	mils := duration.Milliseconds() % 100

	timeText := fmt.Sprintf("El. Time: %02d : %02d : %02d\n", mins, secs, mils)

	rl.DrawText(timeText, offset, getTextPos(2), 12, rl.White)
}
