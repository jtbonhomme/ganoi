package game

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/jtbonhomme/ganoi/internal/towers"
)

// Game is a struct to describe a Hanoi Tower game
type Game struct {
	Game   *tl.Game
	Towers []towers.Tower
	Level  *tl.BaseLevel
}

// New creates a new game
func New(n int) *Game {
	_game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	t := make([]towers.Tower, n)

	game := Game{
		Game:   _game,
		Level:  level,
		Towers: t,
	}

	game.Level.AddEntity(tl.NewText(2, 2, "HANOI TOWERS", tl.ColorWhite, tl.ColorBlack))

	game.Level.AddEntity(tl.NewText(15, 6, "a", tl.ColorWhite, tl.ColorBlack))
	game.Level.AddEntity(tl.NewText(45, 6, "b", tl.ColorWhite, tl.ColorBlack))
	game.Level.AddEntity(tl.NewText(75, 6, "c", tl.ColorWhite, tl.ColorBlack))

	tower1 := tl.NewRectangle(15, 8, 1, 6, tl.ColorWhite)
	tower2 := tl.NewRectangle(45, 8, 1, 6, tl.ColorWhite)
	tower3 := tl.NewRectangle(75, 8, 1, 6, tl.ColorWhite)

	game.Level.AddEntity(tower1)
	game.Level.AddEntity(tower2)
	game.Level.AddEntity(tower3)

	base4 := tl.NewRectangle((15 - (3*4 + 1)), (9 + 4), (2*4+1)*3, 1, tl.ColorYellow)
	base3 := tl.NewRectangle((15 - (3*3 + 1)), (9 + 3), (2*3+1)*3, 1, tl.ColorBlue)
	base2 := tl.NewRectangle((15 - (3*2 + 1)), (9 + 2), (2*2+1)*3, 1, tl.ColorGreen)
	base1 := tl.NewRectangle((15 - (3*1 + 1)), (9 + 1), (2*1+1)*3, 1, tl.ColorRed)
	base0 := tl.NewRectangle((15 - (3*0 + 1)), (9 + 0), (2*0+1)*3, 1, tl.ColorCyan)

	game.Level.AddEntity(base4)
	game.Level.AddEntity(base3)
	game.Level.AddEntity(base2)
	game.Level.AddEntity(base1)
	game.Level.AddEntity(base0)

	game.Game.Screen().SetLevel(game.Level)

	return &game
}

// Start run the game (this is a blocking call)
func (g *Game) Start() {
	g.Game.Start()
}
