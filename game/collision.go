package game

import (
	"math"
)

type Collision struct {
	g *Game
}

func (c *Collision) BallRect(screenWidth int, screenHeight int) {
	ball := c.g.Ball
	rect := c.g.Rect

	closestX := c.closest(ball.X, rect.X, rect.X+rect.Width)
	closestY := c.closest(ball.Y, rect.Y, rect.Y+rect.Height)

	dx := ball.X - closestX
	dy := ball.Y - closestY

	distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	if distance < ball.Radius {
		if distance == 0 {
			distance = 0.01
		}
		nx := dx / distance
		ny := dy / distance

		overlap := ball.Radius - distance
		ball.X += nx * overlap
		ball.Y += ny * overlap
	}
}

func (c *Collision) ColiderRect(screenWidth int, screenHeight int) {
	ball := c.g.Ball
	collider := c.g.Colider
	rect := c.g.Rect

	closestX := c.closest(collider.X, rect.X, rect.X+rect.Width)
	closestY := c.closest(collider.Y, rect.Y, rect.Y+rect.Height)

	dx := collider.X - closestX
	dy := collider.Y - closestY

	distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	if distance < ball.Radius {
		ball.IsOnSurface = true
		ball.Jumps = 2
	} else {
		ball.IsOnSurface = false
	}
}

func (c *Collision) closest(value float32, min float32, max float32) float32 {
	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}
