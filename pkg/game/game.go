package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/louisinger/spaceGoInvader/pkg/entity"
	"github.com/louisinger/spaceGoInvader/pkg/background"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
)

const (
	shipTileHeight = 64
	shipTileWidth = 66
)

// Game implements ebiten.Game interface.
type Game struct{
	background *background.Background
	Entities []entity.Entity
	eventHandler EventHandler
}

// NewGame inits the game instance
func NewGame(screenWidth float64, screenHeight float64) (*Game, error) {
	backgroundImage, _, err := ebitenutil.NewImageFromFile("./assets/space.png")
	if err != nil {
		return nil, err
	}

	game	:= &Game{
		background: background.NewBackground(backgroundImage),
		Entities: make([]entity.Entity, 0),
	}

	game.eventHandler = newEventHandler(game)

	shipImage, _, err := ebitenutil.NewImageFromFile("./assets/ships.png")
	if err != nil {
		return nil, err
	}
	shipImage = shipImage.SubImage(image.Rect(0, 0, shipTileHeight, shipTileWidth)).(*ebiten.Image)

	bulletImage, _, err := ebitenutil.NewImageFromFile("./assets/laser-03.png")
	if err != nil {
		return nil, err
	}

	ship, _ := entity.NewShip(shipImage, bulletImage, 100, 300, screenWidth - shipTileWidth, screenHeight - shipTileHeight)
	game.Entities = append(game.Entities, ship)

	return game, nil
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (game *Game) Update() error {
	log.Print(game.Entities)
	game.background.Update()
	// update entities
	for _, entity := range game.Entities {
		events, err := entity.Update()
		if err != nil {
			return err
		}

		for _, event := range events {
			err := game.eventHandler.handle(event)
			if err != nil {
				return err
			}
		}
	}
  return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (game *Game) Draw(screen *ebiten.Image) {	
	game.background.Draw(screen)
	for _, entity := range game.Entities {
		entity.Draw(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 640, 480
}