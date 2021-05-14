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
	Towers []*towers.Tower
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

	game := Game{
		Game:   _game,
		Level:  level,
		Towers: []*towers.Tower{},
	}

	game.Level.AddEntity(tl.NewText(2, 2, "HANOI TOWERS", tl.ColorWhite, tl.ColorBlack))
	game.Level.AddEntity(tl.NewText(2, 17, "Press CTRL+C to exit", tl.ColorWhite, tl.ColorBlack))

	for i := 0; i < 3; i++ {
		game.Level.AddEntity(tl.NewText(15+(towers.TowerWidth*i), 6, fmt.Sprintf("%d", i), tl.ColorWhite, tl.ColorBlack))
		tower := towers.New(i)
		game.Towers = append(game.Towers, tower)
		game.Level.AddEntity(tower.Rec)
	}

	for i := 4; i >= 0; i-- {
		base := bases.New(i, 0)
		game.Level.AddEntity(base.Rec)
		game.Towers[0].Bases = append(game.Towers[0].Bases, base)
	}
	game.Game.Screen().SetLevel(game.Level)

	return &game
}

// Start run the game (this is a blocking call)
func (g *Game) Start() {
	g.Game.Start()
}

// Move moves top base from a tower to another tower
func (g *Game) Move(from, to int) error {
	if from < 0 || from > 2 || to < 0 || to > 2 {
		return errors.New("incorrect parameters")
	}
	fromHeight := len(g.Towers[from].Bases)
	if fromHeight < 1 {
		return errors.New("no base to move")
	}
	base, err := g.Towers[from].Pop()
	if err != nil {
		return fmt.Errorf("base pop failed because %w", err)
	}
	err = g.Towers[to].Push(base)
	if err != nil {
		return fmt.Errorf("base push failed because %w", err)
	}
	return nil
}
