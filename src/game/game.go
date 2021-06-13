package game

type Game struct {
}

func New() *Game {
	return &Game{}
}

func (g *Game) Draw() {
	drawBoard()
	drawCircle(250, 150)
	drawCross(175, 25)
}

var (
	boardSize float32 = 450
)

func drawBoard() {
	drawLine(getVector2(boardSize/3, 0), getVector2(boardSize/3, boardSize))
	drawLine(getVector2(boardSize*2/3, 0), getVector2(boardSize*2/3, boardSize))

	drawLine(getVector2(0, boardSize/3), getVector2(boardSize, boardSize/3))
	drawLine(getVector2(0, boardSize*2/3), getVector2(boardSize, boardSize*2/3))
}
