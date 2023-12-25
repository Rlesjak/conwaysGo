package cell

import "image"

type Cell struct {
	X          int
	Y          int
	Neighbours int
}

type CellNeighbour int

const (
	Top CellNeighbour = iota
	TopRight
	Right
	BottomRight
	Bottom
	BottomLeft
	Left
	TopLeft
	NeighbourCount
)

func (c *Cell) Equal(cin *Cell) bool {
	return (c.X == cin.X) && (c.Y == cin.Y)
}

func (c *Cell) EqualPos(x int, y int) bool {
	return (c.X == x) && (c.Y == y)
}

// XXX
// XAX
// XXX
func (c *Cell) GetSurroundingCells() [NeighbourCount]image.Point {
	neighbours := [NeighbourCount]image.Point{}

	neighbours[Top] = image.Point{
		X: c.X,
		Y: c.Y - 1,
	}

	neighbours[TopRight] = image.Point{
		X: c.X + 1,
		Y: c.Y - 1,
	}

	neighbours[Right] = image.Point{
		X: c.X + 1,
		Y: c.Y,
	}

	neighbours[BottomRight] = image.Point{
		X: c.X + 1,
		Y: c.Y + 1,
	}

	neighbours[Bottom] = image.Point{
		X: c.X,
		Y: c.Y + 1,
	}

	neighbours[BottomLeft] = image.Point{
		X: c.X - 1,
		Y: c.Y + 1,
	}

	neighbours[Left] = image.Point{
		X: c.X - 1,
		Y: c.Y,
	}

	neighbours[TopLeft] = image.Point{
		X: c.X - 1,
		Y: c.Y - 1,
	}

	return neighbours
}
