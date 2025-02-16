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
		widget.ContainerOpts.BackgroundImage(eimage.NewNineSliceColor(color3)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(30)),
		)),
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
			"Snake Game",
			face,
			color.White,
		),
		widget.TextOpts.Position(
			widget.TextPositionStart,
			widget.TextPositionStart,
		),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
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
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
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
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
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
