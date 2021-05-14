package bases

import (
	tl "github.com/JoelOtter/termloop"
)

// Base is a struct that describes a tower floor
type Base struct {
	Rank  int
	Color tl.Attr
	Rec   *tl.Rectangle
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

func New(rank, pos int) *Base {
	b := Base{
		Rank: rank,
	}
	x := ((15 + 30*pos) - ((3 * b.Rank) + 1))
	y := (9 + b.Rank)
	width := ((2 * b.Rank) + 1) * 3
	b.Rec = tl.NewRectangle(x, y, width, 1, getBaseColor(rank))
	return &b
}
