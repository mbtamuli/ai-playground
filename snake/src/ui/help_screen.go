package ui

import (
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

func helpScreenWindow() *widget.Window {

	// load button text font
	face, _ := loadFont(20)

	// load the font for the window title
	titleFace, _ := loadFont(12)

	// Create the contents of the window
	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{236, 236, 236, 255})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	windowContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("This snake game help text", face, color.Black),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	))

	// Create the titlebar for the window
	titleContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{150, 150, 150, 255})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Help", titleFace, color.NRGBA{254, 255, 255, 255}),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	))

	// Create the new window object. The window object is not tied to a container. Its location and
	// size are set manually using the SetLocation method on the window and added to the UI with ui.AddWindow()
	// Set the Button callback below to see how the window is added to the UI.
	return widget.NewWindow(
		//Set the main contents of the window
		widget.WindowOpts.Contents(windowContainer),
		//Set the titlebar for the window (Optional)
		widget.WindowOpts.TitleBar(titleContainer, 25),
		//Set the window above everything else and block input elsewhere
		widget.WindowOpts.Modal(),
		//Set how to close the window. CLICK_OUT will close the window when clicking anywhere
		//that is not a part of the window object
		widget.WindowOpts.CloseMode(widget.CLICK_OUT),
		//Indicates that the window is draggable. It must have a TitleBar for this to work
		widget.WindowOpts.Draggable(),
		//Set the window resizeable
		widget.WindowOpts.Resizeable(),
		//Set the minimum size the window can be
		widget.WindowOpts.MinSize(200, 100),
		//Set the maximum size a window can be
		widget.WindowOpts.MaxSize(300, 300),
		//Set the callback that triggers when a move is complete
		widget.WindowOpts.MoveHandler(func(args *widget.WindowChangedEventArgs) {
			fmt.Println("Window Moved")
		}),
		//Set the callback that triggers when a resize is complete
		widget.WindowOpts.ResizeHandler(func(args *widget.WindowChangedEventArgs) {
			fmt.Println("Window Resized")
		}),
	)
}
