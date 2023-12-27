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
	aliveMap   AliveMap
	generation int64
}

type CoordsKey struct {
	x int
	y int
}

type AliveMap map[CoordsKey]*cell.Cell

func New() Life {
	return Life{
		generation: 0,
		aliveMap:   make(AliveMap),
	}
}

func (l *Life) GetAlive() *(AliveMap) {
	return &l.aliveMap
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

	start := time.Now()

	// Collect cells to kill, we dont want to kill them
	// while iterating over the map because it would
	// change the map
	cellsToKill := []*cell.Cell{}

	// Collect all dead cells that need to be processed with rule 4
	var deadCells = []cell.Cell{}

	// Process alive cells
	for _, aliveCell := range l.aliveMap {

		if aliveCell.Neighbours < 2 {
			// Rule 1
			cellsToKill = append(cellsToKill, aliveCell)
		} else if aliveCell.Neighbours > 3 {
			// Rule 3
			cellsToKill = append(cellsToKill, aliveCell)
		}

		// Get relevant dead cells
		// Go throug all sourounding cells of a live cell
		// and find only the dead ones
		neigbourCords := aliveCell.GetSurroundingCells()
		for _, ncord := range neigbourCords {
			if l.GetCell(ncord.X, ncord.Y) == nil {
				deadCells = append(deadCells, cell.Cell{
					X: ncord.X,
					Y: ncord.Y,
				})
			}
		}
	}

	deadCellStart := time.Now()

	// Process relevant dead cells
	// Rule 4
	pointChannel := make(chan image.Point)
	var wg sync.WaitGroup

	for _, deadCell := range deadCells {

		wg.Add(1)

		go func(cl cell.Cell) {
			defer wg.Done()

			aliveNeigbourIndexes := l.getNeighbours(&cl)
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

	mutateStart := time.Now()

	// After processing all rules apply all changes at once

	// Spawn cells
	for _, point := range toSpawn {
		l.Spawn(point.X, point.Y)
	}

	// Kill cells
	for _, cellToKill := range cellsToKill {
		l.Kill(cellToKill.X, cellToKill.Y)
	}

	// fmt.Println("Processing dead took: ", time.Since(startNow))

	if cfg.Debug {
		delta := deadCellStart.Sub(start)
		deltaSpawn := mutateStart.Sub(deadCellStart)
		fmt.Println("Generation ", l.generation, " took: ", time.Since(start))
		fmt.Println("rules 123 took: ", delta)
		fmt.Println("rule 4 took: ", time.Since(deadCellStart))
		fmt.Println("mutating took: ", deltaSpawn)
		fmt.Println("----------------------------------------------------------")
	}
}

func (l *Life) Spawn(x int, y int) {

	// Ignore duplicates
	if l.GetCell(x, y) != nil {
		return
	}

	newCell := cell.Cell{
		X:          x,
		Y:          y,
		Neighbours: 0,
	}

	nbrs := l.getNeighbours(&newCell)

	// Count the number of alive neigbours for the new cell
	newCell.Neighbours = len(nbrs)

	// To all neigbours of the new cell add one more neigbour
	for _, z := range nbrs {
		z.Neighbours++
	}

	// Add new cell to the cords map
	l.aliveMap[CoordsKey{
		x: x,
		y: y,
	}] = &newCell

}

func (l *Life) Kill(x int, y int) {
	// find cell in map
	cellToKill := l.GetCell(x, y)
	if cellToKill == nil {
		// Cell does not exist
		return
	}

	nbrs := l.getNeighbours(cellToKill)

	// To all neigbours of the new cell subtract one more neigbour
	for _, z := range nbrs {
		z.Neighbours--
	}

	// Remove cell from cords map
	delete(l.aliveMap, CoordsKey{
		x: x,
		y: y,
	})
}

func (l *Life) GetCell(x int, y int) *cell.Cell {

	i, ok := l.aliveMap[CoordsKey{
		x: x,
		y: y,
	}]

	if ok {
		return i
	}

	return nil
}

func (l *Life) getNeighbours(cell *cell.Cell) (nbrs []*cell.Cell) {
	return getNeighbours(cell, &l.aliveMap)
}

func getNeighbours(cell *cell.Cell, aliveMap *(AliveMap)) (nbrs []*cell.Cell) {
	nbrCoordinates := cell.GetSurroundingCells()

	for _, nbr := range nbrCoordinates {
		if c, ok := (*aliveMap)[CoordsKey{
			x: nbr.X,
			y: nbr.Y,
		}]; ok {
			nbrs = append(nbrs, c)
		}
	}

	return
}

func (l *Life) PrintCells() {
	fmt.Print("#### SAVE ####\n\n")
	for _, cell := range l.aliveMap {
		fmt.Printf("Spawn(%d, %d)\n", cell.X, cell.Y)
	}
	fmt.Println("\n#### END SAVE ####")
}
