package life

import (
	"fmt"

	"github.com/Rlesjak/conwaysGo/cell"
)

type Life struct {
	alive      []cell.Cell
	generation int64
}

func New() Life {
	return Life{
		alive:      []cell.Cell{},
		generation: 0,
	}
}

func (l *Life) GetAlive() *([]cell.Cell) {
	return &l.alive
}

func (l *Life) Tick() {
	l.generation++

	// Collect all cells to which rules need to be applied
	// Tecnichally rules have to be applied to all cells
	// but it is not necessary to apply rules to blank space ...
	//
	// All cells in a 1 thickness ring around each alive cell
	// will be collected for ruling
	//
	// XXX
	// XAX
	// XXX

	// cellsToProcess := []cell.Cell{}
	// for _, c := range l.alive {
	// 	neighbours := c.GetSurroundingCells()
	// 	cellsToProcess = append(cellsToProcess, neighbours[:]...)
	// }
}

func (l *Life) Spawn(x int, y int) {

	// Ignore duplicates
	for _, c := range l.alive {
		if c.EqualPos(x, y) {
			return
		}
	}

	l.alive = append(l.alive, cell.Cell{
		X: x,
		Y: y,
	})
}

func (l *Life) Kill(x int, y int) {
	// find cell in array

	var i int
	var c cell.Cell
	for i, c = range l.alive {
		if c.EqualPos(x, y) {
			break
		}
	}

	l.alive = append(l.alive[:i], l.alive[(i+1):]...)
}

// Retreves the bounds of life
// func (l *Life) GetBounds() Rect {
// 	rect := Rect{
// 		Top:    l.alive[0].Y,
// 		Bottom: l.alive[0].Y,
// 		Left:   l.alive[0].X,
// 		Right:  l.alive[0].X,
// 	}

// 	for _, c := range l.alive {
// 		if c.X < rect.Left {
// 			rect.Left = c.X
// 		}
// 		if c.X > rect.Right {
// 			rect.Right = c.X
// 		}
// 		if c.Y < rect.Top {
// 			rect.Top = c.Y
// 		}
// 		if c.Y > rect.Bottom {
// 			rect.Bottom = c.Y
// 		}

// 	}

// 	return rect
// }

func (l *Life) Debug() {
	fmt.Println(l.alive)
}
