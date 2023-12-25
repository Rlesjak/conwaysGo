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

	// for i := 1; i <= 50; i++ {
	// 	offset := i * 10
	// 	game.Life.Spawn(10+offset, 10)
	// 	game.Life.Spawn(11+offset, 11)
	// 	game.Life.Spawn(12+offset, 11)
	// 	game.Life.Spawn(11+offset, 12)
	// 	game.Life.Spawn(10+offset, 12)
	// }

	// for i := 1; i <= 50; i++ {
	// 	offset := i * 10
	// 	game.Life.Spawn(10, 10+offset)
	// 	game.Life.Spawn(11, 11+offset)
	// 	game.Life.Spawn(12, 11+offset)
	// 	game.Life.Spawn(11, 12+offset)
	// 	game.Life.Spawn(10, 12+offset)
	// }

	// for i := 1; i <= 50; i++ {
	// 	offset := i * 10
	// 	game.Life.Spawn(10+offset, 10+offset)
	// 	game.Life.Spawn(11+offset, 11+offset)
	// 	game.Life.Spawn(12+offset, 11+offset)
	// 	game.Life.Spawn(11+offset, 12+offset)
	// 	game.Life.Spawn(10+offset, 12+offset)
	// }

	for j := 10; j <= 20; j++ {
		for i := 1; i <= 10; i++ {
			offset := i * j
			game.Life.Spawn(10+offset, 10)
			game.Life.Spawn(11+offset, 11)
			game.Life.Spawn(12+offset, 11)
			game.Life.Spawn(11+offset, 12)
			game.Life.Spawn(10+offset, 12)
		}

		for i := 1; i <= 10; i++ {
			offset := i * j
			game.Life.Spawn(10, 10+offset)
			game.Life.Spawn(11, 11+offset)
			game.Life.Spawn(12, 11+offset)
			game.Life.Spawn(11, 12+offset)
			game.Life.Spawn(10, 12+offset)
		}
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
