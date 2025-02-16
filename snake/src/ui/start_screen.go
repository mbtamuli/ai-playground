package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/mbtamuli/ai-playground/snake/src/ui/resources"
)

// createStartScreen creates the start screen UI
func createStartScreen(manager *UIManager) *ebitenui.UI {
	ui := &ebitenui.UI{}

	rootContainer := createRootContainer()
	title := createTitle(manager.Fonts)
	buttonContainer := createButtonContainer(manager)

	rootContainer.AddChild(title)
	rootContainer.AddChild(buttonContainer)
	ui.Container = rootContainer

	return ui
}

// createRootContainer creates the main container for the start screen
func createRootContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eimage.NewNineSliceColor(gray)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 20, Bottom: 20,
				Left: 20, Right: 20,
			}),
		)),
	)
}

// createTitle creates the game title widget
func createTitle(fonts *resources.Fonts) *widget.Text {
	return widget.NewText(
		widget.TextOpts.Text("Snake Game", fonts.TitleFace, color.Black),
		widget.TextOpts.Position(widget.TextPositionStart, widget.TextPositionStart),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
}

// createButtonContainer creates a container with the game buttons
func createButtonContainer(manager *UIManager) *widget.Container {
	buttonContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 20, Bottom: 20,
			}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)

	startButton := createGameButton(manager.ButtonResources, "Start Game", func(args *widget.ButtonClickedEventArgs) {
		//log.Println("Start Game button clicked")
		manager.ShowGameScreen()
	})

	helpButton := createGameButton(manager.ButtonResources, "Help", func(args *widget.ButtonClickedEventArgs) {
		manager.ShowHelpScreen()
	})

	buttonContainer.AddChild(startButton)
	buttonContainer.AddChild(helpButton)
	return buttonContainer
}

// createGameButton creates a button with standard styling
func createGameButton(button *resources.ButtonResources, text string, handler func(*widget.ButtonClickedEventArgs)) *widget.Button {
	buttonTextColor := &widget.ButtonTextColor{
		Idle:    color.Black,
		Pressed: color.Black,
	}

	return widget.NewButton(
		widget.ButtonOpts.Text(text, button.Face, buttonTextColor),
		widget.ButtonOpts.Image(button.Image),
		widget.ButtonOpts.ClickedHandler(handler),
		widget.ButtonOpts.TextPadding(widget.Insets{Left: 20, Right: 20}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)
}
