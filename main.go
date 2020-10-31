package main

import (
	_ "image/png"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	g "github.com/louisinger/spaceGoInvader/pkg/game"
)

func main() {
		game, err := g.NewGame(640, 480)
		if err != nil {
			log.Fatal(err)
		}
    // Sepcify the window size as you like. Here, a doulbed size is specified.
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Your game's title")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}