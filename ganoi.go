package ganoi

import (
	"time"

	"github.com/jtbonhomme/ganoi/internal/game"
)

func hanoi(g *game.Game) {

}

func Run() {
	g := game.New(4)

	go func(g *game.Game) {
		var err error
		time.Sleep(time.Second * 2)
		err = g.Move(0, 2)
		if err != nil {
			return
		}
		time.Sleep(time.Second * 2)
		err = g.Move(0, 1)
		if err != nil {
			return
		}
		time.Sleep(time.Second * 2)
		err = g.Move(2, 1)
		if err != nil {
			return
		}
	}(g)

	g.Start()
}
