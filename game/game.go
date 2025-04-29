package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Ball *Ball
	Rect *Rect
}

func NewGame() *Game {
	ball := NewBall(100, 50, 20, 10, 0.5, 0.1, color.White)

	return &Game{
		Ball: ball,
		Rect: &Rect{
			Height: 100,
			Width:  100,
			Color:  color.RGBA{225, 0, 0, 225},
		},
	}
}

func (g *Game) Update() error {
	w, h := ebiten.WindowSize()
	c := Collision{g: g}
	c.BallRect(w, h)
	g.Ball.Update(w, h)
	g.Rect.Update(w, h)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.Rect.Draw(screen)
	g.Ball.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
