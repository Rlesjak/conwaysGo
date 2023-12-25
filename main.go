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

	// Gosper glider gun
	game.Life.Spawn(6, 20)
	game.Life.Spawn(6, 21)
	game.Life.Spawn(7, 20)
	game.Life.Spawn(7, 21)
	game.Life.Spawn(16, 21)
	game.Life.Spawn(16, 20)
	game.Life.Spawn(16, 22)
	game.Life.Spawn(17, 19)
	game.Life.Spawn(18, 18)
	game.Life.Spawn(19, 18)
	game.Life.Spawn(17, 23)
	game.Life.Spawn(18, 24)
	game.Life.Spawn(19, 24)
	game.Life.Spawn(20, 21)
	game.Life.Spawn(21, 19)
	game.Life.Spawn(21, 23)
	game.Life.Spawn(22, 20)
	game.Life.Spawn(22, 21)
	game.Life.Spawn(22, 22)
	game.Life.Spawn(23, 21)
	game.Life.Spawn(26, 20)
	game.Life.Spawn(26, 19)
	game.Life.Spawn(26, 18)
	game.Life.Spawn(27, 20)
	game.Life.Spawn(27, 19)
	game.Life.Spawn(27, 18)
	game.Life.Spawn(28, 17)
	game.Life.Spawn(28, 21)
	game.Life.Spawn(30, 17)
	game.Life.Spawn(30, 16)
	game.Life.Spawn(30, 21)
	game.Life.Spawn(30, 22)
	game.Life.Spawn(40, 18)
	game.Life.Spawn(40, 19)
	game.Life.Spawn(41, 18)
	game.Life.Spawn(41, 19)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
