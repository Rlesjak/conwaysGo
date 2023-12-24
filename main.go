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
	game.Life.Spawn(10, 20)
	game.Life.Spawn(10, 21)
	game.Life.Spawn(20, 20)
	game.Life.Spawn(21, 20)
	game.Life.Spawn(22, 20)
	game.Life.Spawn(23, 20)
	game.Life.Spawn(23, 21)
	game.Life.Spawn(23, 22)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
