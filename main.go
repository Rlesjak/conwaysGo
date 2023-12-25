package main

import (
	"log"

	"github.com/Rlesjak/conwaysGo/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowPosition(0, 0)
	ebiten.SetWindowSize(1200, 1000)
	ebiten.SetWindowTitle("Conways Life")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(500)

	game := game.New()
	game.Life.Spawn(10, 10)
	game.Life.Spawn(11, 11)
	game.Life.Spawn(12, 11)
	game.Life.Spawn(11, 12)
	game.Life.Spawn(10, 12)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
