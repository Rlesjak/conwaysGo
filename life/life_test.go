package life

import (
	"testing"

	"github.com/Rlesjak/conwaysGo/cell"
)

func TestSingleCellDeath(t *testing.T) {
	lifeInstance := New()
	lifeInstance.Spawn(1, 1)
	lifeInstance.Evolve()
	if lifeInstance.GetCell(1, 1) != nil {
		t.Error("Single cell did not die!")
	}
}

func TestUnderPopulation(t *testing.T) {
	lifeInstance := New()

	lifeInstance.Spawn(1, 1)
	lifeInstance.Spawn(1, 2)

	lifeInstance.Spawn(4, 1)
	lifeInstance.Spawn(4, 2)

	lifeInstance.Evolve()

	alive := lifeInstance.GetAlive()
	if len(*alive) != 0 {
		t.Error("All cells did not die!")
	}
}

func TestCellBirth(t *testing.T) {

	l := New()

	l.Spawn(2, 0)
	l.Spawn(3, 1)
	l.Spawn(1, 2)

	l.Evolve()

	// 2,1 should now be alive, all other should be dead
	shouldBeDead := []*cell.Cell{
		l.GetCell(2, 0),
		l.GetCell(3, 1),
		l.GetCell(1, 2),
	}

	for _, res := range shouldBeDead {
		if res != nil {
			t.Error("Cell that should be dead found alive!")
		}
	}

	if l.GetCell(2, 1) == nil {
		t.Error("Cell that should be alive found dead!")
	}

}

func TestBlinker(t *testing.T) {
	l := New()

	// XXXXX
	// XOOOX
	// XXXXX
	l.Spawn(1, 1)
	l.Spawn(2, 1)
	l.Spawn(3, 1)

	l.Evolve()

	// Should now be
	// XXOXX
	// XXOXX
	// XXOXX
	shouldBeFound := []*cell.Cell{
		l.GetCell(2, 0),
		l.GetCell(2, 1),
		l.GetCell(2, 2),
	}

	shouldBeDead := []*cell.Cell{
		l.GetCell(1, 1),
		l.GetCell(3, 1),
	}

	for _, res := range shouldBeFound {
		if res == nil {
			t.Error("Cell that should be alive found dead!")
		}
	}

	for _, res := range shouldBeDead {
		if res != nil {
			t.Error("Cell that should be dead found alive!")
		}
	}

	l.Evolve()

	// Should now be bac to
	// XXXXX
	// XOOOX
	// XXXXX
	shouldBeFound = []*cell.Cell{
		l.GetCell(1, 1),
		l.GetCell(2, 1),
		l.GetCell(3, 1),
	}

	shouldBeDead = []*cell.Cell{
		l.GetCell(2, 0),
		l.GetCell(2, 2),
	}

	for _, res := range shouldBeFound {
		if res == nil {
			t.Error("Cell that should be alive found dead!")
		}
	}

	for _, res := range shouldBeDead {
		if res != nil {
			t.Error("Cell that should be dead found alive!")
		}
	}
}

func TestGetCell(t *testing.T) {
	l := New()

	l.Spawn(1, 2)

	if l.GetCell(1, 2) == nil {
		t.Error("Cell not found!")
	}

	l.Kill(1, 2)

	if l.GetCell(1, 2) != nil {
		t.Error("Cell found!")
	}
}
