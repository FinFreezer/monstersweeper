package main

import (
	"log"

	ms "github.com/finfreezer/monstersweeper/monstersweeper"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := ms.MsweeperGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(ms.SCREENWIDTH, ms.SCREENHEIGHT)
	ebiten.SetWindowTitle("Monstersweeper")
	ebiten.SetTPS(15)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
