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

	startStopButton Button
	clearButton     Button
	clearCallback   func()
}

var backgroundColor = imcolor.RGBA{0x0f, 0x0f, 0x0f, 0x88}

func NewGameMenu() GameMenu {

	startStopButton := NewButton(
		10,
		50,
		100,
		40,
		color.Green,
		"START",
	)

	clearButton := NewButton(
		120,
		50,
		60,
		40,
		color.Red,
		"CLEAR",
	)

	return GameMenu{
		Width:           200,
		Height:          100,
		IsRunning:       false,
		startStopButton: startStopButton,
		clearButton:     clearButton,
	}
}

func (m *GameMenu) SetClearCallback(callback func()) {
	m.clearCallback = callback
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

	// Draw buttons

	m.clearButton.Draw(dst)

	m.startStopButton.SetContent(color.Green, "START")
	if m.IsRunning {
		m.startStopButton.SetContent(color.Red, "PAUSE")
	}
	m.startStopButton.Draw(dst)
}

func (m *GameMenu) UpdateState(generation int) {
	m.generation = generation
}

func (m *GameMenu) CaptureMouseClick(mx, my int) (captured bool) {

	if (mx <= m.Width) && (my <= m.Height) {

		// Play pause button position
		if m.startStopButton.IsTarget(mx, my) {
			m.IsRunning = !m.IsRunning
		} else if m.clearButton.IsTarget(mx, my) && m.clearCallback != nil {
			m.clearCallback()
		}

		return true
	}
	return false
}
