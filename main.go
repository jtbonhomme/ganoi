package main

import (
    tl "github.com/JoelOtter/termloop"
)

func main() {
    game := tl.NewGame()
    game.Screen().SetFps(60)

    level := tl.NewBaseLevel(tl.Cell{
        Bg: tl.ColorBlack,
        Fg: tl.ColorWhite,
        Ch: ' ',
    })

    tower1 := tl.NewRectangle(27, 5, 1, 6, tl.ColorWhite)
    tower2 := tl.NewRectangle(52, 5, 1, 6, tl.ColorWhite)
    tower3 := tl.NewRectangle(77, 5, 1, 6, tl.ColorWhite)

    base4 := tl.NewRectangle(17, 10, 21, 1, tl.ColorBlue)
    base3 := tl.NewRectangle(20, 9, 15, 1, tl.ColorGreen)
    base2 := tl.NewRectangle(23, 8, 9, 1, tl.ColorRed)
    base1 := tl.NewRectangle(26, 7, 3, 1, tl.ColorCyan)

    level.AddEntity(tl.NewText(2, 2, "HANOI TOWERS", tl.ColorWhite, tl.ColorBlack))

    level.AddEntity(tower1)
    level.AddEntity(tower2)
    level.AddEntity(tower3)

    level.AddEntity(base4)
    level.AddEntity(base3)
    level.AddEntity(base2)
    level.AddEntity(base1)

    game.Screen().SetLevel(level)
/*    go func() {
        for i := 0 ; i < 10; i++ {
            base4.SetPosition(10, 10+i)
        }
    }()*/
    game.Start()
}
