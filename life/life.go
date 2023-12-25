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

	// Copy the current alive cells
	prevGen := append(make([]cell.Cell, 0, len(l.alive)), l.alive...)

	var deadCells = []cell.Cell{}

	// Process alive cells
	for _, aliveCell := range prevGen {

		if aliveCell.Neighbours < 2 {
			// Rule 1
			l.Kill(aliveCell.X, aliveCell.Y)
		} else if aliveCell.Neighbours > 3 {
			// Rule 3
			l.Kill(aliveCell.X, aliveCell.Y)
		}

		// Get relevant dead cells
		// Go throug all sourounding cells of a live cell
		// and find only the dead ones
		neigbourCords := aliveCell.GetSurroundingCells()
		for _, cord := range neigbourCords {
			if l.cordsToIndex(cord.X, cord.Y) == -1 {
				deadCells = append(deadCells, cell.Cell{
					X: cord.X,
					Y: cord.Y,
				})
			}
		}
	}

	// Process relevant dead cells
	// Rule 4
	for _, deadCell := range deadCells {
		aliveNeigbourIndexes := getNeighbourIndexes(&deadCell, &prevGen)
		if len(aliveNeigbourIndexes) == 3 {
			l.Spawn(deadCell.X, deadCell.Y)
		}
	}
}

func (l *Life) Spawn(x int, y int) {

	// Ignore duplicates
	for _, c := range l.alive {
		if c.EqualPos(x, y) {
			return
		}
	}

	newCell := cell.Cell{
		X:          x,
		Y:          y,
		Neighbours: 0,
	}

	nbrIndexes := l.getNeighbourIndexes(&newCell)

	// Count the number of alive neigbours for the new cell
	newCell.Neighbours = len(nbrIndexes)

	// To all neigbours of the new cell add one more neigbour
	for _, z := range nbrIndexes {
		l.alive[z].Neighbours++
	}

	// Add new cell to the alive array
	l.alive = append(l.alive, newCell)
}

func (l *Life) Kill(x int, y int) {
	// find cell in array

	cellIndex := l.cordsToIndex(x, y)
	if cellIndex == -1 {
		// Cell not alive
		return
	}

	cellToKill := l.alive[cellIndex]
	nbrIndexes := l.getNeighbourIndexes(&cellToKill)

	// To all neigbours of the new cell subtract one more neigbour
	for _, z := range nbrIndexes {
		l.alive[z].Neighbours--
	}

	l.alive = append(l.alive[:cellIndex], l.alive[(cellIndex+1):]...)
}

func (l *Life) cordsToIndex(x int, y int) int {
	for i, c := range l.alive {
		if c.EqualPos(x, y) {
			return i
		}
	}

	return -1
}

func (l *Life) getNeighbourIndexes(cell *cell.Cell) (nbrIndexes []int) {
	return getNeighbourIndexes(cell, &l.alive)
}

func getNeighbourIndexes(cell *cell.Cell, aliveCells *([]cell.Cell)) (nbrIndexes []int) {
	nbrCoordinates := cell.GetSurroundingCells()

	for i, c := range *aliveCells {
		// Find indexes (in l.alive slice) of all alive neigbours of the newCell
		for _, nbr := range nbrCoordinates {
			if c.EqualPos(nbr.X, nbr.Y) {
				nbrIndexes = append(nbrIndexes, i)
			}
		}
	}

	return
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