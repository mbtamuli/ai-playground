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
	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eimage.NewNineSliceColor(color3)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Direction(widget.DirectionVertical))),
	)

	buttonImage, err := loadButtonImage()
	if err != nil {
		log.Fatal(err)
	}

	face, err := loadFont(20)
	if err != nil {
		log.Fatal(err)
	}

	title := widget.NewText(
		widget.TextOpts.Text(
			"Welcome to Snake Game - AI playground",
			face,
			color.Black,
		),
	)

	// Create a "Start game" button
	startButton := widget.NewButton(
		widget.ButtonOpts.Text("Start Game", face, &widget.ButtonTextColor{
			Idle:    color2,
			Hover:   color4,
			Pressed: color1,
		}),
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			log.Println("Start Game button clicked")
		}),
	)

	// Create a "Start game" button
	helpButton := widget.NewButton(
		widget.ButtonOpts.Text("Help", face, &widget.ButtonTextColor{
			Idle:    color2,
			Hover:   color4,
			Pressed: color1,
		}),
		widget.ButtonOpts.Image(buttonImage),
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
	)

	rootContainer.AddChild(title)
	rootContainer.AddChild(startButton)
	rootContainer.AddChild(helpButton)

	ui.Container = rootContainer

	return ui
}
