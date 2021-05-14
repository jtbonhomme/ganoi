package bases

import (
	tl "github.com/JoelOtter/termloop"
)

// Base is a struct that describes a tower floor
type Base struct {
	towersX []int
	Rank    int
	Pos     int
	X       int
	Y       int
	Width   int
	Height  int
	Level   int
	Color   tl.Attr
	Rec     *tl.Rectangle
}

func getBaseColors() []tl.Attr {
	return []tl.Attr{
		tl.ColorCyan,
		tl.ColorRed,
		tl.ColorGreen,
		tl.ColorBlue,
		tl.ColorYellow,
	}
}

func getBaseColor(rank int) tl.Attr {
	c := getBaseColors()
	return c[rank]
}

func New(rank, pos, level int) *Base {
	b := Base{
		towersX: []int{15, 45, 75},
		Rank:    rank,
		Pos:     pos,
		Level:   level,
		Height:  1,
	}
	b.X = (b.towersX[b.Pos] - ((3 * b.Rank) + 1))
	b.Y = (9 + b.Rank)
	b.Width = ((2 * b.Rank) + 1) * 3
	b.Rec = tl.NewRectangle(b.X, b.Y, b.Width, b.Height, getBaseColor(rank))
	return &b
}
