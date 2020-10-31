package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Entity represents all the things appearing in the game.
type Entity interface {
	Update() error
	Draw(screen *ebiten.Image)
}

