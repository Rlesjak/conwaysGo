package life

import (
	"testing"
)

func TestSingleCellDeath(t *testing.T) {
	lifeInstance := New()
	lifeInstance.Spawn(1, 1)
	lifeInstance.Evolve()
	if lifeInstance.CordsToIndex(1, 1) != -1 {
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
	shouldBeDead := []int{
		l.CordsToIndex(2, 0),
		l.CordsToIndex(3, 1),
		l.CordsToIndex(1, 2),
	}

	for _, res := range shouldBeDead {
		if res != -1 {
			t.Error("Cell that should be dead found alive!")
		}
	}

	if l.CordsToIndex(2, 1) == -1 {
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
	shouldBeFound := []int{
		l.CordsToIndex(2, 0),
		l.CordsToIndex(2, 1),
		l.CordsToIndex(2, 2),
	}

	shouldBeDead := []int{
		l.CordsToIndex(1, 1),
		l.CordsToIndex(3, 1),
	}

	for _, res := range shouldBeFound {
		if res == -1 {
			t.Error("Cell that should be alive found dead!")
		}
	}

	for _, res := range shouldBeDead {
		if res != -1 {
			t.Error("Cell that should be dead found alive!")
		}
	}

	l.Evolve()

	// Should now be bac to
	// XXXXX
	// XOOOX
	// XXXXX
	shouldBeFound = []int{
		l.CordsToIndex(1, 1),
		l.CordsToIndex(2, 1),
		l.CordsToIndex(3, 1),
	}

	shouldBeDead = []int{
		l.CordsToIndex(2, 0),
		l.CordsToIndex(2, 2),
	}

	for _, res := range shouldBeFound {
		if res == -1 {
			t.Error("Cell that should be alive found dead!")
		}
	}

	for _, res := range shouldBeDead {
		if res != -1 {
			t.Error("Cell that should be dead found alive!")
		}
	}
}
