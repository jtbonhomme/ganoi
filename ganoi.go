package ganoi

import (
	"github.com/jtbonhomme/ganoi/internal/game"
)

func Run() {
	g := game.New(4)

	go func() {
		g.Move(0, 2)
	}()
	g.Start()
}
