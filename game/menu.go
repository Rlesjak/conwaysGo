package game

import (
	"fmt"
	imcolor "image/color"

	"github.com/Rlesjak/conwaysGo/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameMenu struct {
	Width  int
	Height int

	generation int
	IsRunning  bool
}

var backgroundColor = imcolor.RGBA{0x0f, 0x0f, 0x0f, 0x88}

func NewGameMenu() GameMenu {
	return GameMenu{
		Width:     200,
		Height:    100,
		IsRunning: false,
	}
}

func (m *GameMenu) Draw(dst *ebiten.Image) {
	vector.DrawFilledRect(
		dst,
		0,
		0,
		float32(m.Width),
		float32(m.Height),
		backgroundColor,
		false,
	)

	fpsAndTpsString := fmt.Sprintf("FPS: %d, TPS: %d", int(ebiten.ActualFPS()), int(ebiten.ActualTPS()))

	ebitenutil.DebugPrintAt(
		dst,
		fpsAndTpsString,
		10,
		10,
	)

	ebitenutil.DebugPrintAt(
		dst,
		fmt.Sprintf("Generation: %d", m.generation),
		10,
		30,
	)

	// Draw play pause button

	buttonColor := color.Green
	buttonText := "START"
	if m.IsRunning {
		buttonColor = color.Red
		buttonText = "PAUSE"
	}

	vector.DrawFilledRect(
		dst,
		10,
		50,
		150,
		40,
		buttonColor,
		false,
	)

	ebitenutil.DebugPrintAt(
		dst,
		buttonText,
		70,
		63,
	)
}

func (m *GameMenu) UpdateState(generation int) {
	m.generation = generation
}

func (m *GameMenu) CaptureMouseClick(mx, my int) (captured bool) {

	if (mx <= m.Width) && (my <= m.Height) {

		// Play pause button position
		if (mx > 10) && (mx < 10+150) &&
			(my > 50) && (my < 50+40) {
			m.IsRunning = !m.IsRunning
		}

		return true
	}
	return false
}
