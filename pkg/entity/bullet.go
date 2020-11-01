package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

// Bullet if the entity of bullet fired by ships
type Bullet struct {
	image *ebiten.Image
	position Position
	dY float64
	dX float64
	rotationAngle float64
	alive bool
}

// BulletFactory is a func that generate a bullet at x, y
type BulletFactory func(x float64, y float64) *Bullet

// NewBulletFactory returns a factory function
func NewBulletFactory(bulletSprite *ebiten.Image, dX float64, dY float64, screenWidth float64, screenHeight float64) BulletFactory {
	return func (x float64, y float64) *Bullet {
		return &Bullet{
			image: bulletSprite,
			dY: dY,
			dX: dX,
			position: Position{
				X: x, 
				Y: y,
				// max greater than screen in order to the bullet disapear
				maxX: screenWidth + 64,
				maxY: screenHeight + 64,
			},
			rotationAngle: rotationAngle(dX, dY),
			alive: true,
		}
	}
}

// Update the state of the bullet
func (b *Bullet) Update() ([]Event, error) {
	events := make([]Event, 0)

	isOut := b.position.translate(b.dX, b.dY)
	
	if isOut {
		events = append(events, &RemoveEntityEvent{
			BaseEvent: BaseEvent{ source: b },
			EntityToRemove: b,
		})
	}

	return events, nil
}

// Draw the bullet on the screen
func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(b.position.X, b.position.Y)
	// op.GeoM.Rotate(b.rotationAngle)
	screen.DrawImage(b.image, op)
}

func rotationAngle(dX float64, dY float64) float64 {
	alpha := math.Atan(dX / dY)
	angle := math.Pi / 2 - alpha
	return angle
}