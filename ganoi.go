package ganoi

import (
	"time"

	"github.com/jtbonhomme/ganoi/internal/game"
)

func Run() {
	g := game.New(4)

	go func() {
		time.Sleep(time.Second * 2)
		g.Move(0, 2)
		time.Sleep(time.Second * 2)
		g.Move(0, 1)
		time.Sleep(time.Second * 2)
		g.Move(2, 1)
	}()
	g.Start()
}
