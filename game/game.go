package game

import (
	"image"
	"math"

	"github.com/Rlesjak/conwaysGo/color"
	"github.com/Rlesjak/conwaysGo/geometry"
	"github.com/Rlesjak/conwaysGo/grid"
	"github.com/Rlesjak/conwaysGo/life"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	grid grid.Grid
	Life life.Life

	dragStartPos *image.Point
}

func New() Game {
	return Game{
		grid: grid.Grid{
			CellSize: 25,
			Camera: geometry.Rect{
				X:      0,
				Y:      0,
				Width:  0,
				Height: 0,
			},
		},
		Life: life.New(),
	}
}

func (g *Game) pan() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle) {
		mX, mY := ebiten.CursorPosition()
		g.dragStartPos = &image.Point{
			X: mX + g.grid.Camera.X,
			Y: mY + g.grid.Camera.Y,
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonMiddle) {
		g.dragStartPos = nil
	}

	if g.dragStartPos != nil {
		mX, mY := ebiten.CursorPosition()
		g.grid.Camera.X = g.dragStartPos.X - mX
		g.grid.Camera.Y = g.dragStartPos.Y - mY
	}
}

func (g *Game) zoom() {
	mX, mY := ebiten.CursorPosition()

	oldGX, oldGY := g.grid.ViewportToGridCords(mX, mY)

	_, scroll := ebiten.Wheel()
	g.grid.CellSize += float32(scroll)
	g.grid.CellSize = float32(math.Max(5, math.Min(100, float64(g.grid.CellSize))))

	newGX, newGY := g.grid.ViewportToGridCords(mX, mY)

	g.grid.Camera.X += int(float32(oldGX-newGX) * g.grid.CellSize)
	g.grid.Camera.Y += int(float32(oldGY-newGY) * g.grid.CellSize)
}

func (g *Game) Update() error {

	// Spawning cells on click
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mX, mY := ebiten.CursorPosition()
		gX, gY := g.grid.ViewportToGridDescreteCords(mX, mY)
		g.Life.Spawn(gX, gY)
	}

	// Handle panning
	g.pan()

	// Handle zooming
	g.zoom()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screenBounds := screen.Bounds()
	g.grid.Camera.Width = screenBounds.Dx()
	g.grid.Camera.Height = screenBounds.Dy()

	// Print clock
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("Clock: %d", g.clock))
	g.grid.Draw(screen, g.Life.GetAlive())

	mX, mY := ebiten.CursorPosition()

	vector.DrawFilledCircle(
		screen,
		float32(mX),
		float32(mY),
		4,
		color.Green,
		false,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
