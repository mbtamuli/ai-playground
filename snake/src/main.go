package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mbtamuli/ai-playground/snake/src/game"
	"github.com/mbtamuli/ai-playground/snake/src/ui"
)

// Game represents the main game state
type SnakePlayground struct {
	uiManager *ui.UIManager
}

// Update handles the game logic updates
func (g SnakePlayground) Update() error {
	g.uiManager.Update()
	return nil
}

// Draw renders the game screen
func (g SnakePlayground) Draw(screen *ebiten.Image) {
	g.uiManager.Draw(screen)
}

// Layout returns the game's screen dimensions
func (g SnakePlayground) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
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

	gameManager := game.NewGameManager()
	game := &SnakePlayground{
		uiManager: ui.NewUIManager(gameManager),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
