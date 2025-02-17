package resources

import (
	"github.com/ebitenui/ebitenui/image"
)

const (
	background = "assets/PNG/Grey/Default/button_square_flat.png"
)

type BackgroundResources struct {
	Image *image.NineSlice
}

func GetBackgroundResource() (*BackgroundResources, error) {
	image, err := createScalableNineSliceImage(background, 64/2, 64/2)
	if err != nil {
		return nil, err
	}

	return &BackgroundResources{
		Image: image,
	}, nil
}
