package resources

import (
	"github.com/ebitenui/ebitenui/widget"

	"github.com/mbtamuli/ai-playground/snake/assets"
)

const (
	slideHandle          = "PNG/Blue/Default/slide_hangle.png"
	slideHorizontalColor = "PNG/Blue/Default/slide_horizontal_color.png"
)

type SliderResources struct {
	TrackImage  *widget.SliderTrackImage
	HandleImage *widget.ButtonImage
}

func CreateSliderResources() (*SliderResources, error) {
	track, err := assets.CreateScalableNineSliceImage(slideHorizontalColor, 96/2, 16/2)
	if err != nil {
		return nil, err
	}

	handle, err := assets.CreateScalableNineSliceImage(slideHandle, 24, 32)
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
