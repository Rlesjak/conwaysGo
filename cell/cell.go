package cell

type Cell struct {
	X int
	Y int
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
	return c.X == cin.X && c.Y == cin.Y
}

func (c *Cell) EqualPos(x int, y int) bool {
	return c.X == x && c.Y == x
}

// XXX
// XAX
// XXX
func (c *Cell) GetSurroundingCells() [NeighbourCount]Cell {
	neighbours := [NeighbourCount]Cell{}

	neighbours[Top] = Cell{
		X: c.X,
		Y: c.Y - 1,
	}

	neighbours[TopRight] = Cell{
		X: c.X + 1,
		Y: c.Y - 1,
	}

	neighbours[Right] = Cell{
		X: c.X + 1,
		Y: c.Y,
	}

	neighbours[BottomRight] = Cell{
		X: c.X + 1,
		Y: c.Y + 1,
	}

	neighbours[Bottom] = Cell{
		X: c.X,
		Y: c.Y + 1,
	}

	neighbours[BottomLeft] = Cell{
		X: c.X - 1,
		Y: c.Y + 1,
	}

	neighbours[Left] = Cell{
		X: c.X - 1,
		Y: c.Y,
	}

	neighbours[TopLeft] = Cell{
		X: c.X - 1,
		Y: c.Y - 1,
	}

	return neighbours
}
