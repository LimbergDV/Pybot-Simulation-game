package main

import (
	"log"

	"pybot-simulator/config"
	"pybot-simulator/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Robot Recolector")
	
	g := game.NewGame(config.ScreenWidth, config.ScreenHeight)
	
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}