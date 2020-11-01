package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
	"log"
)

type direction int

// Ship is the player
type Ship struct {
	image    *ebiten.Image
	position Position
	xSpeed float64
	ySpeed float64
	acceleration float64
	bulletFactory BulletFactory
	reloadTime time.Duration
	isAbleToFire bool
}

// NewShip build and create a ship using sprite and a initial position (x,y)
func NewShip(sprite *ebiten.Image, bulletSprite *ebiten.Image, x float64, y float64, boundWidth float64, boundHeight float64) (*Ship, error) {
	ship := &Ship{
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
		bulletFactory: NewBulletFactory(bulletSprite, 0, -15, boundWidth, boundHeight),
		reloadTime: time.Duration(800),
		isAbleToFire: true,
	}
	return ship, nil
}

func (s *Ship) fireBullet() Event {
	if s.isAbleToFire {
		s.isAbleToFire = false

		log.Print("fire bullet")
		bullet := s.bulletFactory(s.position.X + 28, s.position.Y)
		go func () {
			time.Sleep(s.reloadTime * time.Millisecond)
			s.isAbleToFire = true
		} ()
		return &AddEntityEvent{ BaseEvent: BaseEvent{ source: s }, EntityToAdd: bullet }
	}
	return nil
}

// Update the state of the Ship
func (s *Ship) Update() ([]Event, error) {
	events := make([]Event, 0)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fireEvent := s.fireBullet()
		if fireEvent != nil {
			events = append(events, fireEvent)
		}
	}

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

	return events, nil
}

// Draw the ship
func (s *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.position.X, s.position.Y)
	screen.DrawImage(s.image, op)
}
