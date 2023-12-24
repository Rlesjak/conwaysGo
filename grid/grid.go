package grid

import (
	"math"

	"github.com/Rlesjak/conwaysGo/cell"
	"github.com/Rlesjak/conwaysGo/color"
	"github.com/Rlesjak/conwaysGo/geometry"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Grid struct {
	CellSize float32
	Camera   geometry.Rect
}

func (g *Grid) Draw(dst *ebiten.Image, cells *([]cell.Cell)) {
	g.drawEmptyGrid(dst)

	for _, cell := range *cells {
		g.drawFilledCell(dst, int(cell.X), int(cell.Y))
	}
}

func (g *Grid) ViewportToGridDescreteCords(viewportX int, viewportY int) (x, y int) {
	x = int(float64(g.Camera.X+viewportX) / float64(g.CellSize))
	y = int(float64(g.Camera.Y+viewportY) / float64(g.CellSize))
	return
}

func (g *Grid) ViewportToGridCords(viewportX int, viewportY int) (x, y float64) {
	x = float64(g.Camera.X+viewportX) / float64(g.CellSize)
	y = float64(g.Camera.Y+viewportY) / float64(g.CellSize)
	return
}

func (g *Grid) getBorderPadding() float32 {
	return float32(math.Max(0.5, float64(g.CellSize/50)))
}

func (g *Grid) GridDescreteToViewPortCords(gridDescreteX int, gridDescreteY int) (x, y float32) {
	gridAbsX := float32(gridDescreteX) * g.CellSize
	gridAbsY := float32(gridDescreteY) * g.CellSize

	x = gridAbsX - float32(g.Camera.X)
	y = gridAbsY - float32(g.Camera.Y)
	return
}

// Converts the window coordinate system bounds to the grid coordinate system bounds
// Converts from viewport in pixels to view port in grid coordinates
func (g *Grid) getVisibleGridBounds() geometry.Rect {

	// Find the x coordinate of the top intersecting tiles
	x := float64(g.Camera.X) / float64(g.CellSize)
	x = math.Floor(x)

	// Find the y coordinate of the left intersecting tiles
	y := float64(g.Camera.Y) / float64(g.CellSize)
	y = math.Floor(y)

	return geometry.Rect{
		X:      int(x),
		Y:      int(y),
		Width:  int(math.Ceil(float64(g.Camera.Width) / float64(g.CellSize))),
		Height: int(math.Ceil(float64(g.Camera.Height) / float64(g.CellSize))),
	}
}

func (g *Grid) drawEmptyCell(dst *ebiten.Image, gridDescreteX int, gridDescreteY int) {

	padding := g.getBorderPadding()

	screenX, screenY := g.GridDescreteToViewPortCords(gridDescreteX, gridDescreteY)

	vector.DrawFilledRect(
		dst,
		screenX,
		screenY,
		g.CellSize,
		g.CellSize,
		color.Gray,
		false,
	)

	vector.DrawFilledRect(
		dst,
		screenX+padding,
		screenY+padding,
		g.CellSize-2*padding,
		g.CellSize-2*padding,
		color.LightGray,
		false,
	)
}

func (g *Grid) drawFilledCell(dst *ebiten.Image, gridX int, gridY int) {

	screenX, screenY := g.GridDescreteToViewPortCords(gridX, gridY)

	vector.DrawFilledRect(
		dst,
		screenX,
		screenY,
		g.CellSize,
		g.CellSize,
		color.Black,
		false,
	)
}

func (g *Grid) drawEmptyGrid(dst *ebiten.Image) {

	grid0X, grid0Y := g.GridDescreteToViewPortCords(0, 0)

	defer vector.DrawFilledCircle(
		dst,
		grid0X,
		grid0Y,
		5,
		color.Red,
		false,
	)

	visibleGridBounds := g.getVisibleGridBounds()

	for i := 0; i <= visibleGridBounds.Width; i++ {
		for j := 0; j <= visibleGridBounds.Height; j++ {
			g.drawEmptyCell(dst, i+visibleGridBounds.X, j+visibleGridBounds.Y)
		}
	}
}
