package grid

import (
	"log"
	"math"

	"github.com/Rlesjak/conwaysGo/cell"
	"github.com/Rlesjak/conwaysGo/color"
	"github.com/Rlesjak/conwaysGo/geometry"
	"github.com/Rlesjak/conwaysGo/life"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

type Grid struct {
	CellSize float32
	Camera   geometry.Rect

	font *sfnt.Font
}

func New() Grid {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	return Grid{
		CellSize: 25,
		Camera: geometry.Rect{
			X:      0,
			Y:      0,
			Width:  0,
			Height: 0,
		},
		font: tt,
	}
}

func (g *Grid) Draw(dst *ebiten.Image, cells *(life.AliveMap)) {
	g.drawEmptyGrid(dst)

	// Draw only the cells that are visible
	visibleGridBounds := g.getVisibleGridBounds()

	for i := 0; i < visibleGridBounds.Width; i++ {
		for j := 0; j < visibleGridBounds.Height; j++ {
			cell := (*cells)[life.CoordsKey{
				X: visibleGridBounds.X + i,
				Y: visibleGridBounds.Y + j,
			}]
			if cell != nil {
				g.drawFilledCell(dst, cell)
			}
		}
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

// Converts the grid coordinate system coordinates to the viewport coordinate system coordinates
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

func (g *Grid) drawFilledCell(dst *ebiten.Image, cell *cell.Cell) {

	screenX, screenY := g.GridDescreteToViewPortCords(cell.X, cell.Y)

	vector.DrawFilledRect(
		dst,
		screenX,
		screenY,
		g.CellSize,
		g.CellSize,
		color.Black,
		false,
	)

	// fontSize := float64(g.CellSize * 2 / 3)

	// mplusNormalFont, _ := opentype.NewFace(g.font, &opentype.FaceOptions{
	// 	Size:    fontSize,
	// 	DPI:     72,
	// 	Hinting: font.HintingVertical,
	// })

	// text.Draw(
	// 	dst,
	// 	fmt.Sprintf("%d", cell.Neighbours),
	// 	mplusNormalFont,
	// 	int(screenX)+int(fontSize/2),
	// 	int(screenY)+int(fontSize*1.1),
	// 	color.Gray,
	// )
}

func (g *Grid) drawEmptyGrid(dst *ebiten.Image) {

	// Draw background
	dst.Fill(color.LightGray)

	visibleGridBounds := g.getVisibleGridBounds()
	padding := g.getBorderPadding()
	dblPadding := 2 * padding

	// Draw vertical lines
	for i := 0; i <= visibleGridBounds.Width; i++ {
		screenX, screenY := g.GridDescreteToViewPortCords(visibleGridBounds.X+i, visibleGridBounds.Y)
		vector.DrawFilledRect(
			dst,
			screenX-padding,
			screenY,
			dblPadding,
			float32(visibleGridBounds.Height)*g.CellSize,
			color.Gray,
			false,
		)
	}

	// Draw horizontal lines
	for j := 0; j <= visibleGridBounds.Height; j++ {
		screenX, screenY := g.GridDescreteToViewPortCords(visibleGridBounds.X, visibleGridBounds.Y+j)
		vector.DrawFilledRect(
			dst,
			screenX,
			screenY-padding,
			float32(visibleGridBounds.Width)*g.CellSize,
			dblPadding,
			color.Gray,
			false,
		)
	}

	// Get the coordinates of the grid ORIGIN in the viewport
	grid0X, grid0Y := g.GridDescreteToViewPortCords(0, 0)

	vector.DrawFilledCircle(
		dst,
		grid0X,
		grid0Y,
		5,
		color.Red,
		false,
	)
}
