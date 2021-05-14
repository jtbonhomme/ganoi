package main

import (
	"github.com/jtbonhomme/ganoi/internal/game"
)

func main() {
	g := game.New(4)

	/*    go func() {
	      for i := 0 ; i < 10; i++ {
	          base4.SetPosition(10, 10+i)
	      }
	  }()*/
	g.Start()
}
