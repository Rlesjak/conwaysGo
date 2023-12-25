package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	rect  image.Rectangle
	color color.Color
	label string
}

func NewButton(x, y, width, height int, color color.Color, label string) Button {
	return Button{
		rect: image.Rectangle{
			Min: image.Point{
				X: x,
				Y: y,
			},
			Max: image.Point{
				X: x + width,
				Y: y + height,
			},
		},
		color: color,
		label: label,
	}
}

func (b *Button) SetContent(color color.Color, label string) {
	b.color = color
	b.label = label
}

func (b *Button) IsTarget(x, y int) bool {
	if (x > b.rect.Min.X) && (x < b.rect.Max.X) &&
		(y > b.rect.Min.Y) && (y < b.rect.Max.Y) {
		return true
	}

	return false
}

func (b *Button) Draw(dst *ebiten.Image) {

	vector.DrawFilledRect(
		dst,
		float32(b.rect.Min.X),
		float32(b.rect.Min.Y),
		float32(b.rect.Dx()),
		float32(b.rect.Dy()),
		b.color,
		false,
	)

	ebitenutil.DebugPrintAt(
		dst,
		b.label,
		b.rect.Min.X+(b.rect.Dx()/2)-len(b.label)*3,
		b.rect.Min.Y+(b.rect.Dy()/2)-4,
	)
}
