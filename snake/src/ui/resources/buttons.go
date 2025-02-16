package resources

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	idleButton    = "assets/PNG/Extra/Default/button_rectangle_line.png"
	pressedButton = "assets/PNG/Extra/Default/button_rectangle_depth_line.png"
)

// ButtonResources contains all the necessary resources for rendering a button
type ButtonResources struct {
	Image   *widget.ButtonImage
	Text    *widget.ButtonTextColor
	Face    text.Face
	Padding widget.Insets
}

// CreateButtonResources creates a new button resources instance with the
// specified fonts and returns all necessary components for button rendering
func CreateButtonResources(fontSet *Fonts) (*ButtonResources, error) {
	idleButtonImage, err := CreateScalableNineSliceImage(idleButton, 12, 0)
	if err != nil {
		return nil, err
	}

	pressedButtonImage, err := CreateScalableNineSliceImage(pressedButton, 12, 0)
	if err != nil {
		return nil, err
	}

	buttonImages := &widget.ButtonImage{
		Idle:    idleButtonImage,
		Pressed: pressedButtonImage,
	}

	return &ButtonResources{
		Image: buttonImages,
		Text: &widget.ButtonTextColor{
			Idle:    color.Black,
			Pressed: color.Black,
		},
		Face: fontSet.Face,
		Padding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}
