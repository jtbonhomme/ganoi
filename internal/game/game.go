package game

import (
	"errors"
	"fmt"

	tl "github.com/JoelOtter/termloop"
	"github.com/jtbonhomme/ganoi/internal/bases"
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

	for i := 0; i < 3; i++ {
		game.Level.AddEntity(tl.NewText(15+(towers.TowerWidth*i), 6, fmt.Sprintf("%d", i), tl.ColorWhite, tl.ColorBlack))
		tower := tl.NewRectangle(15+(towers.TowerWidth*i), 8, 1, 6, tl.ColorWhite)
		game.Level.AddEntity(tower)
	}

	for i := 4; i >= 0; i-- {
		base := bases.New(i, 0, 4-i)
		game.Level.AddEntity(base.Rec)
		game.Towers[0].Bases = append(game.Towers[0].Bases, i)
	}
	game.Game.Screen().SetLevel(game.Level)

	return &game
}

// Start run the game (this is a blocking call)
func (g *Game) Start() {
	g.Game.Start()
}

// Move moves bases from a tower to another tower
func (g *Game) Move(from, to int) error {
	if from < 0 || from > 2 || to < 0 || to > 2 {
		return errors.New("incorrect parameters")
	}

	return nil
}
