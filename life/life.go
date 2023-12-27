package life

import (
	"fmt"
	"image"
	"sync"
	"time"

	"github.com/Rlesjak/conwaysGo/cell"
	"github.com/Rlesjak/conwaysGo/cfg"
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

func (l *Life) GetGeneration() int {
	return int(l.generation)
}

func (l *Life) Clear() {
	l.PrintCells()
	*l = New()
}

func (l *Life) Evolve() {
	l.generation++

	startNow := time.Now()

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
			if l.CordsToIndex(cord.X, cord.Y) == -1 {
				deadCells = append(deadCells, cell.Cell{
					X: cord.X,
					Y: cord.Y,
				})
			}
		}
	}

	// Process relevant dead cells
	// Rule 4
	pointChannel := make(chan image.Point)
	var wg sync.WaitGroup

	for _, deadCell := range deadCells {

		wg.Add(1)

		go func(cl cell.Cell) {
			defer wg.Done()

			aliveNeigbourIndexes := getNeighbourIndexes(&cl, &prevGen)
			if len(aliveNeigbourIndexes) == 3 {
				pointChannel <- image.Point{
					X: cl.X,
					Y: cl.Y,
				}
			}
		}(deadCell)
	}

	go func() {
		wg.Wait()
		close(pointChannel)
	}()

	toSpawn := []image.Point{}
	for point := range pointChannel {
		toSpawn = append(toSpawn, point)
	}

	for _, point := range toSpawn {
		l.Spawn(point.X, point.Y)
	}

	// fmt.Println("Processing dead took: ", time.Since(startNow))

	if cfg.Debug {
		fmt.Println("Generation ", l.generation, " took: ", time.Since(startNow))
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

	cellIndex := l.CordsToIndex(x, y)
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

func (l *Life) CordsToIndex(x int, y int) int {
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

func (l *Life) PrintCells() {
	fmt.Print("#### SAVE ####\n\n")
	for _, cell := range l.alive {
		fmt.Printf("Spawn(%d, %d)\n", cell.X, cell.Y)
	}
	fmt.Println("\n#### END SAVE ####")
}
