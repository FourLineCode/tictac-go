package game

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	textureColor rl.Color = rl.Black
	lineWidth    float32  = 10
	circleRadius float32  = boardSize / 9
	crossWidth   float32  = circleRadius * 2
)

func drawLine(start, end rl.Vector2) {
	rl.DrawLineEx(start, end, lineWidth, textureColor)
}

func drawCircle(x, y int32) {
	rl.DrawCircle(x, y, circleRadius, rl.Black)
	rl.DrawCircle(x, y, circleRadius-lineWidth, rl.DarkGray)
}

func drawCross(startX, startY float32) {
	drawLine(getVector2(startX, startY), getVector2(startX+crossWidth, startY+crossWidth))
	drawLine(getVector2(startX+crossWidth, startY), getVector2(startX, startY+crossWidth))
}
