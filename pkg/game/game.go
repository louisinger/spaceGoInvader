package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/louisinger/spaceGoInvader/pkg/entity"
	"github.com/louisinger/spaceGoInvader/pkg/background"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
)

const (
	shipTileHeight = 64
	shipTileWidth = 66
)

// Game implements ebiten.Game interface.
type Game struct{
	background *background.Background
	entities []entity.Entity
}

// NewGame inits the game instance
func NewGame(screenWidth float64, screenHeight float64) (*Game, error) {
	backgroundImage, _, err := ebitenutil.NewImageFromFile("./assets/space.png")
	if err != nil {
		return nil, err
	}

	game:= &Game{
		background: background.NewBackground(backgroundImage),
		entities: make([]entity.Entity, 0),
	}

	shipImage, _, err := ebitenutil.NewImageFromFile("./assets/ships.png")
	if err != nil {
		return nil, err
	}
	shipImage = shipImage.SubImage(image.Rect(0, 0, shipTileHeight, shipTileWidth)).(*ebiten.Image)

	ship, _ := entity.NewShip(shipImage, 100, 300, screenWidth - shipTileWidth, screenHeight - shipTileHeight)
	game.entities = append(game.entities, ship)


	return game, nil
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (game *Game) Update() error {
	game.background.Update()
	// update entities
	for _, entity := range game.entities {
		err := entity.Update()
		return err
	}
  return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (game *Game) Draw(screen *ebiten.Image) {	
	game.background.Draw(screen)
	for _, entity := range game.entities {
		entity.Draw(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 640, 480
}