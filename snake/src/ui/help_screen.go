package ui

import (
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"

	"github.com/mbtamuli/ai-playground/snake/src/ui/resources"
)

// createHelpScreen creates the help screen UI
func createHelpScreen(manager *UIManager) *ebitenui.UI {
	ui := &ebitenui.UI{}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(gray)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 20, Bottom: 20,
				Left: 20, Right: 20,
			}),
		)),
	)

	helpContent := createHelpContent(manager.Fonts)
	backButton := createGameButton(manager.ButtonResources, "Back to Menu", func(args *widget.ButtonClickedEventArgs) {
		manager.ShowStartScreen()
	})

	rootContainer.AddChild(createHelpTitleBar(manager.Fonts))
	rootContainer.AddChild(helpContent)
	rootContainer.AddChild(backButton)

	ui.Container = rootContainer
	return ui
}

// createHelpContent creates the main content container for the help window
func createHelpContent(fonts *resources.Fonts) *widget.Container {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(gray)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 20, Bottom: 20,
				Left: 20, Right: 20,
			}),
		)),
	)

	instructions := getInstructions()
	for _, instruction := range instructions {
		container.AddChild(createInstructionText(instruction, fonts))
	}

	return container
}

// getInstructions returns the list of game instructions
func getInstructions() []string {
	return []string{
		"How to Play:",
		"- Use arrow keys to control the snake",
		"- Eat food to grow longer",
		"- Avoid hitting walls and yourself",
		"- Press ESC to pause the game",
		"",
	}
}

// createInstructionText creates a text widget for a single instruction
func createInstructionText(instruction string, fonts *resources.Fonts) *widget.Text {
	return widget.NewText(
		widget.TextOpts.Text(instruction, fonts.Face, color.Black),
	)
}

// createHelpTitleBar creates the title bar for the help window
func createHelpTitleBar(fonts *resources.Fonts) *widget.Container {
	titleContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{150, 150, 150, 255})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	titleContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Game Instructions", fonts.TitleFace, color.Black),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	))

	return titleContainer
}

// createHelpWindow creates the window with the specified content and title bar
func createHelpWindow(content *widget.Container, titleBar *widget.Container) *widget.Window {
	return widget.NewWindow(
		widget.WindowOpts.Contents(content),
		widget.WindowOpts.TitleBar(titleBar, 30),
		widget.WindowOpts.Modal(),
		widget.WindowOpts.CloseMode(widget.CLICK_OUT),
		widget.WindowOpts.Draggable(),
		widget.WindowOpts.Resizeable(),
		widget.WindowOpts.MinSize(300, 400),
		widget.WindowOpts.MaxSize(400, 500),
		widget.WindowOpts.MoveHandler(func(args *widget.WindowChangedEventArgs) {
			fmt.Println("Window Moved")
		}),
		widget.WindowOpts.ResizeHandler(func(args *widget.WindowChangedEventArgs) {
			fmt.Println("Window Resized")
		}),
	)
}
