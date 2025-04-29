package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Colider struct {
	X      float32
	Y      float32
	Father *Ball
	Radius float32
	Color  color.Color
}

func NewColider(x, y, radius float32, color color.Color, father *Ball) *Colider {
	return &Colider{
		X:      x,
		Y:      y,
		Father: father,
		Radius: radius,
		Color:  color,
	}
}

func (c *Colider) Update() {
	c.X = c.Father.X
	c.Y = c.Father.Y + c.Father.Radius
}

func (c *Colider) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, c.X, c.Y, c.Radius, c.Color, true)
}
