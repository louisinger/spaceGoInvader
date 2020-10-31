package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type direction int

// Ship is the player
type Ship struct {
	image    *ebiten.Image
	position Position
	xSpeed float64
	ySpeed float64
	acceleration float64
}

// NewShip build and create a ship using sprite and a initial position (x,y)
func NewShip(sprite *ebiten.Image, x float64, y float64, boundWidth float64, boundHeight float64) (*Ship, error) {
	return &Ship{
		image: sprite,
		position: Position{
			X: x,
			Y: y,
			maxX: boundWidth,
			maxY: boundHeight,
		},
		xSpeed: 0,
		ySpeed: 0,
		acceleration: 0.1,
	}, nil
}

// Update the state of the Ship
func (s *Ship) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if s.ySpeed > 0 {
			s.ySpeed -= 2
		}
		s.ySpeed -= s.acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if s.ySpeed < 0 {
			s.ySpeed += 2
		}
		s.ySpeed += s.acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if s.xSpeed > 0 {
			s.xSpeed -= 2
		}
		s.xSpeed -= s.acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if s.xSpeed < 0 {
			s.xSpeed += 2
		}
		s.xSpeed += s.acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyX) {
		s.xSpeed = 0
		s.ySpeed = 0
	}

	s.position.translate(s.xSpeed, s.ySpeed)

	return nil
}

// Draw the ship
func (s *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.position.X, s.position.Y)
	screen.DrawImage(s.image, op)
}
