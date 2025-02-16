package ui

import (
	"image"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

func StartScreen() *ebitenui.UI {
	ui := &ebitenui.UI{}

	rootContainer := widget.NewContainer(
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

	fonts, err := loadFonts()
	if err != nil {
		log.Fatal(err)
	}

	button, err := newButtonResources(fonts)
	if err != nil {
		log.Fatal(err)
	}

	title := widget.NewText(
		widget.TextOpts.Text(
			"Snake Game",
			fonts.titleFace,
			color.Black,
		),
		widget.TextOpts.Position(
			widget.TextPositionStart,
			widget.TextPositionStart,
		),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

	buttonTextColor := &widget.ButtonTextColor{
		Idle:    color.Black,
		Pressed: color.Black,
	}

	// Create a "Start game" button
	startButton := widget.NewButton(
		widget.ButtonOpts.Text("Start Game", button.face, buttonTextColor),
		widget.ButtonOpts.Image(button.image),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			log.Println("Start Game button clicked")
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{Left: 20, Right: 20}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)

	// Create a "Start game" button
	helpButton := widget.NewButton(
		widget.ButtonOpts.Text("Help", button.face, buttonTextColor),
		widget.ButtonOpts.Image(button.image),
		widget.ButtonOpts.TextPadding(widget.Insets{Left: 20, Right: 20}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			window := helpScreenWindow()
			//Get the preferred size of the content
			x, y := window.Contents.PreferredSize()
			//Create a rect with the preferred size of the content
			r := image.Rect(0, 0, x, y)
			//Use the Add method to move the window to the specified point
			r = r.Add(image.Point{100, 50})
			//Set the windows location to the rect.
			window.SetLocation(r)

			ui.AddWindow(window)
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)

	// Create a centered container for buttons
	buttonContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)

	rootContainer.AddChild(title)
	buttonContainer.AddChild(startButton)
	buttonContainer.AddChild(helpButton)
	rootContainer.AddChild(buttonContainer)

	ui.Container = rootContainer

	return ui
}
