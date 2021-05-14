package ganoi

import (
	"github.com/jtbonhomme/ganoi/internal/game"
)

func Run() {
	g := game.New(4)

	go func() {
		g.Move(0, 2)
		//	      for i := 0 ; i < 10; i++ {
		//	          base4.SetPosition(10, 10+i)
		//	      }
	}()
	g.Start()
}
