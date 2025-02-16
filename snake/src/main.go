package main

import (
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mbtamuli/ai-playground/snake/src/ui"
)

type Game struct {
	ui *ebitenui.UI
}

func (g *Game) Update() error {
	g.ui.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		ui: ui.StartScreen(),
	}); err != nil {
		log.Fatal(err)
	}
}
