package main

import (
    tl "github.com/JoelOtter/termloop"
)

type Tower struct {
    Name string
    Pos int
    Height int
}

type Game struct {
    Game *tl.Game
    Towers []Tower
    Level  *tl.BaseLevel
}

func New(n int) *Game {
    _game := tl.NewGame()
    level := tl.NewBaseLevel(tl.Cell{
        Bg: tl.ColorBlack,
        Fg: tl.ColorWhite,
        Ch: ' ',
    })
    towers := make([]Tower, n)

    game := Game{
        Game: _game,
        Level: level,
        Towers: towers,
    }

    tower1 := tl.NewRectangle(27, 5, 1, 6, tl.ColorWhite)
    tower2 := tl.NewRectangle(52, 5, 1, 6, tl.ColorWhite)
    tower3 := tl.NewRectangle(77, 5, 1, 6, tl.ColorWhite)

    game.Level.AddEntity(tl.NewText(2, 2, "HANOI TOWERS", tl.ColorWhite, tl.ColorBlack))

    game.Level.AddEntity(tower1)
    game.Level.AddEntity(tower2)
    game.Level.AddEntity(tower3)

    base4 := tl.NewRectangle(17, 10, 21, 1, tl.ColorBlue)
    base3 := tl.NewRectangle(20, 9, 15, 1, tl.ColorGreen)
    base2 := tl.NewRectangle(23, 8, 9, 1, tl.ColorRed)
    base1 := tl.NewRectangle(26, 7, 3, 1, tl.ColorCyan)

    game.Level.AddEntity(base4)
    game.Level.AddEntity(base3)
    game.Level.AddEntity(base2)
    game.Level.AddEntity(base1)

    game.Game.Screen().SetLevel(game.Level)

    return &game
}

func (g *Game) Start() {
    g.Game.Start()
}

func main() {
    game := New(4)

/*    go func() {
        for i := 0 ; i < 10; i++ {
            base4.SetPosition(10, 10+i)
        }
    }()*/
    game.Start()
}
