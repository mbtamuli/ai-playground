package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	gray = color.NRGBA{238, 244, 247, 255} // #eef4f7
)

// UIManager handles the UI state and screen switching
type UIManager struct {
	currentUI *ebitenui.UI
}

// NewUIManager creates a new UI manager instance
func NewUIManager() *UIManager {
	manager := &UIManager{}
	manager.ShowStartScreen()
	return manager
}

// ShowStartScreen switches to the start screen
func (m *UIManager) ShowStartScreen() {
	m.currentUI = createStartScreen(m)
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
