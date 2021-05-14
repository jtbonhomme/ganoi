package towers

import (
	"errors"
	"fmt"

	tl "github.com/JoelOtter/termloop"
	"github.com/jtbonhomme/ganoi/internal/bases"
)

const (
	TowerHeight int = 6
	TowerWidth  int = 30
)

// Tower is a struct to describe a tower in Hanoi Tower game
type Tower struct {
	Pos   int
	Bases []*bases.Base
	Rec   *tl.Rectangle
}

func New(i int) *Tower {
	return &Tower{
		Pos:   i,
		Bases: []*bases.Base{},
		Rec:   tl.NewRectangle(15+(TowerWidth*i), 8, 1, 6, tl.ColorWhite),
	}
}

// Pop removes a base from the tower
func (t *Tower) Pop() (*bases.Base, error) {
	l := len(t.Bases)
	b := t.Bases[l-1]
	t.Bases = t.Bases[:l-1]
	return b, nil
}

// Rank returns the rank of the last base pushed on the tower
func (t *Tower) Rank() int {
	l := len(t.Bases)
	if l > 0 {
		return t.Bases[l-1].Rank
	}
	return 0
}

// Push stacks a base on the tower
func (t *Tower) Push(b *bases.Base) error {
	if b == nil {
		return errors.New("cannot push a nil base")
	}
	l := len(t.Bases)

	rank := t.Rank()

	if l > 0 && rank < b.Rank {
		return fmt.Errorf("impossible to push a base of rank %d on last base (rank %d)", b.Rank, rank)
	}
	b.Rec.SetPosition(15+(t.Pos*30)-((3*b.Rank)+1), (13 - l))
	t.Bases = append(t.Bases, b)
	return nil
}
