package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mbtamuli/ai-playground/snake/src/ui/resources"
)

var (
	gray = color.NRGBA{238, 244, 247, 255} // #eef4f7
)

// UIManager handles the UI state and screen switching
type UIManager struct {
	currentUI *ebitenui.UI
	//gameManager         *game.GameManager
	Fonts               *resources.Fonts
	ButtonResources     *resources.ButtonResources
	SliderResources     *resources.SliderResources
	BackgroundResources *resources.BackgroundResources
}

// NewUIManager creates a new UI manager instance
func NewUIManager() *UIManager {
	fonts, err := resources.SetupFontSystem()
	if err != nil {
		panic(err)
	}

	button, err := resources.CreateButtonResources(fonts)
	if err != nil {
		panic(err)
	}

	slider, err := resources.CreateSliderResources()
	if err != nil {
		panic(err)
	}

	background, err := resources.GetBackgroundResource()
	if err != nil {
		panic(err)
	}

	manager := &UIManager{
		//gameManager:         gameManager,
		Fonts:               fonts,
		ButtonResources:     button,
		SliderResources:     slider,
		BackgroundResources: background,
	}
	manager.ShowStartScreen()
	return manager
}

// ShowStartScreen switches to the start screen
func (m *UIManager) ShowStartScreen() {
	m.currentUI = createStartScreen(m)
}

// ShowStartScreen switches to the start screen
func (m *UIManager) ShowGameScreen() {
	m.currentUI = createGameScreen(m)
}

// ShowHelpScreen switches to the help screen
func (m *UIManager) ShowHelpScreen() {
	m.currentUI = createHelpScreen(m)
}

// Update updates the current UI
func (m *UIManager) Update() {
	if m.currentUI != nil {
		m.currentUI.Update()
	}
}

// Draw draws the current UI
func (m *UIManager) Draw(screen *ebiten.Image) {
	if m.currentUI != nil {
		m.currentUI.Draw(screen)
	}
}
