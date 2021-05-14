package bases

import (
	tl "github.com/JoelOtter/termloop"
)

// Base is a struct that describes a tower floor
type Base struct {
	Rank   int
	Pos    int
	X      int
	Y      int
	Width  int
	Height int
	Level  int
	Color  tl.Attr
}

func New(rank, pos, level int, color tl.Attr) *Base {
	b := Base{
		Rank:   rank,
		Pos:    pos,
		Level:  level,
		Color:  color,
		X:      0,
		Y:      0,
		Width:  0,
		Height: 1,
	}

	return &b
}
