package resources

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/mbtamuli/ai-playground/snake/assets"
)

const (
	fontFaceRegular = "Font/kenney-future.ttf"
	fontFaceNarrow  = "Font/kenney-future-narrow.ttf"
)

// Fonts holds different text.Face instances for various text rendering needs
type Fonts struct {
	Face      text.Face
	TitleFace text.Face
}

// SetupFontSystem initializes and returns the font collection needed for the UI
func SetupFontSystem() (*Fonts, error) {
	defaultFont, err := assets.InitFontFace(fontFaceRegular, 12)
	if err != nil {
		return nil, err
	}

	headerFont, err := assets.InitFontFace(fontFaceNarrow, 24)
	if err != nil {
		return nil, err
	}

	return &Fonts{
		Face:      defaultFont,
		TitleFace: headerFont,
	}, nil
}
