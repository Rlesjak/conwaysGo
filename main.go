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

	game := game.New()
	game.Life.Summon(10, 10)
	game.Life.Summon(10, 20)
	game.Life.Summon(10, 21)
	game.Life.Summon(20, 20)
	game.Life.Summon(21, 20)
	game.Life.Summon(22, 20)
	game.Life.Summon(23, 20)
	game.Life.Summon(23, 21)
	game.Life.Summon(23, 22)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
