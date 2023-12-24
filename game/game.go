package game

import (
	"github.com/Rlesjak/conwaysGo/geometry"
	"github.com/Rlesjak/conwaysGo/grid"
	"github.com/Rlesjak/conwaysGo/life"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid grid.Grid
	Life life.Life
}

func New() Game {
	return Game{
		grid: grid.Grid{
			CellSize: 25,
			Camera:   geometry.Rect{},
		},
		Life: life.New(),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screenBounds := screen.Bounds()
	g.grid.Camera.X = -120
	g.grid.Camera.Y = -210
	g.grid.Camera.Width = screenBounds.Dx()
	g.grid.Camera.Height = screenBounds.Dy()

	// Print clock
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("Clock: %d", g.clock))
	g.grid.Draw(screen, g.Life.GetAlive())
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
