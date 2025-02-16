package resources

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	slideHandle                         = "assets/PNG/Blue/Default/slide_hangle.png"
	slide_horizontal_color_section_wide = "assets/PNG/Blue/Default/slide_horizontal_color_section_wide.png"
	slide_horizontal_color_section      = "assets/PNG/Blue/Default/slide_horizontal_color_section.png"
	slideHorizontalColor                = "assets/PNG/Blue/Default/slide_horizontal_color.png"
	slide_horizontal_grey_section_wide  = "assets/PNG/Blue/Default/slide_horizontal_grey_section_wide.png"
	slide_horizontal_grey_section       = "assets/PNG/Blue/Default/slide_horizontal_grey_section.png"
	slide_horizontal_grey               = "assets/PNG/Blue/Default/slide_horizontal_grey.png"
	slide_vertical_color_section_wide   = "assets/PNG/Blue/Default/slide_vertical_color_section_wide.png"
	slide_vertical_color_section        = "assets/PNG/Blue/Default/slide_vertical_color_section.png"
	slide_vertical_color                = "assets/PNG/Blue/Default/slide_vertical_color.png"
	slide_vertical_grey_section_wide    = "assets/PNG/Blue/Default/slide_vertical_grey_section_wide.png"
	slide_vertical_grey_section         = "assets/PNG/Blue/Default/slide_vertical_grey_section.png"
	slide_vertical_grey                 = "assets/PNG/Blue/Default/slide_vertical_grey.png"
)

type SliderResources struct {
	TrackImage  *ebiten.Image
	HandleImage *ebiten.Image
}

func CreateSliderResources(fontSet *Fonts) (*SliderResources, error) {
	track, err := loadEbitenImageFromAsset(slideHorizontalColor)
	if err != nil {
		return nil, err
	}

	handle, err := loadEbitenImageFromAsset(slideHandle)
	if err != nil {
		return nil, err
	}

	return &SliderResources{
		TrackImage:  track,
		HandleImage: handle,
	}, nil
}
