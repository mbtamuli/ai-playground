package resources

import (
	"github.com/ebitenui/ebitenui/widget"
)

const (
	slideHandle          = "assets/PNG/Blue/Default/slide_hangle.png"
	slideHorizontalColor = "assets/PNG/Blue/Default/slide_horizontal_color.png"
)

type SliderResources struct {
	TrackImage  *widget.SliderTrackImage
	HandleImage *widget.ButtonImage
}

func CreateSliderResources() (*SliderResources, error) {
	track, err := createScalableNineSliceImage(slideHorizontalColor, 96/2, 16/2)
	if err != nil {
		return nil, err
	}

	handle, err := createScalableNineSliceImageHandle(slideHandle)
	if err != nil {
		return nil, err
	}

	return &SliderResources{
		TrackImage: &widget.SliderTrackImage{
			Idle: track,
		},
		HandleImage: &widget.ButtonImage{
			Idle:    handle,
			Pressed: handle,
		},
	}, nil
}
