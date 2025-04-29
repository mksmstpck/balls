package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mksmstpck/balls/game"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hairy Ball")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
