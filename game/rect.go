package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Color  color.Color
	Speed  float32
}

func (r *Rect) Update(screenWidth int, screenHeight int) {
	r.X = float32(screenWidth) / 4
	r.Y = float32(screenHeight) - r.Height
}

func (r *Rect) Draw(screen *ebiten.Image) {
	r.Width = float32(screen.Bounds().Dx() / 2)
	r.Height = float32(screen.Bounds().Dy() / 10)
	vector.DrawFilledRect(screen, r.X, r.Y, r.Width, r.Height, r.Color, true)
}
