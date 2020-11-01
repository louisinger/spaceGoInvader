package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Entity represents all the things appearing in the game.
type Entity interface {
	Update() ([]Event, error)
	Draw(screen *ebiten.Image)
}

