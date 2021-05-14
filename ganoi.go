package ganoi

import (
	"github.com/jtbonhomme/ganoi/internal/game"
)

const TowerHeight int = 5

func Run() {
	g := game.New(TowerHeight)

	go func(g *game.Game) {
		err := g.Move(0, 2, TowerHeight)
		if err != nil {
			return
		}

	}(g)

	g.Start()
}
