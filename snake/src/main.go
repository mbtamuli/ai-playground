package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mbtamuli/ai-playground/snake/src/ui"
)

// Game represents the main game state
type Game struct {
	uiManager *ui.UIManager
}

// Update handles the game logic updates
func (g *Game) Update() error {
	g.uiManager.Update()
	return nil
}

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.uiManager.Draw(screen)
}

// Layout returns the game's screen dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	const (
		defaultWindowWidth  = 640
		defaultWindowHeight = 480
		windowTitle         = "Snake Game"
	)

	ebiten.SetWindowSize(defaultWindowWidth, defaultWindowHeight)
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := &Game{
		uiManager: ui.NewUIManager(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
