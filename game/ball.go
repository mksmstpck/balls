package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	accelerationTime time.Duration = 500 * time.Millisecond
	decelerationTime time.Duration = 250 * time.Millisecond
	startSpeed                     = 4.0
)

type Ball struct {
	X               float32
	Y               float32
	Radius          float32
	Color           color.Color
	SpeedX          float32
	SpeedY          float32
	StartSpeed      float32
	MaxSpeed        float32
	Acceleration    float32
	Deceleration    float32
	LastAccelerated time.Time
	LastDecelerated time.Time
}

func (b *Ball) Update(screenWidth int, screenHeight int) {
	// acceleration
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if time.Since(b.LastAccelerated) > accelerationTime {
			if b.SpeedX < b.MaxSpeed {
				b.SpeedX += b.Acceleration
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if time.Since(b.LastAccelerated) > accelerationTime {
			if b.SpeedX > -b.MaxSpeed {
				b.SpeedX -= b.Acceleration
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if time.Since(b.LastAccelerated) > accelerationTime {
			if b.SpeedY < b.MaxSpeed {
				b.SpeedY += b.Acceleration
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if time.Since(b.LastAccelerated) > accelerationTime {
			if b.SpeedY > -b.MaxSpeed {
				b.SpeedY -= b.Acceleration
			}
		}
	}

	//deceleration x
	if !ebiten.IsKeyPressed(ebiten.KeyD) && !ebiten.IsKeyPressed(ebiten.KeyA) {
		if b.SpeedX > 0 {
			b.SpeedX -= b.Deceleration
			if b.SpeedX < 0 {
				b.SpeedX = 0
			}
		} else if b.SpeedX < 0 {
			b.SpeedX += b.Deceleration
			if b.SpeedX > 0 {
				b.SpeedX = 0
			}
		}
	}

	if !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		if b.SpeedY > 0 {
			b.SpeedY -= b.Deceleration
			if b.SpeedY < 0 {
				b.SpeedY = 0
			}
		} else if b.SpeedY < 0 {
			b.SpeedY += b.Deceleration
			if b.SpeedY > 0 {
				b.SpeedY = 0
			}
		}
	}

	b.X += b.SpeedX
	b.Y += b.SpeedY

	// screen edge collision
	if b.X-b.Radius < 0 {
		b.X = b.Radius
	}
	if b.Y-b.Radius < 0 {
		b.Y = b.Radius
	}
	if b.X+b.Radius > float32(screenWidth) {
		b.X = float32(screenWidth) - b.Radius
	}
	if b.Y+b.Radius > float32(screenHeight) {
		b.Y = float32(screenHeight) - b.Radius
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.X, b.Y, b.Radius, b.Color, true)
}
