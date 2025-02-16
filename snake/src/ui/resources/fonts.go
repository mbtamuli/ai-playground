package resources

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	fontFaceRegular = "assets/Font/kenney-future.ttf"
	fontFaceNarrow  = "assets/Font/kenney-future-narrow.ttf"
)

// Fonts holds different text.Face instances for various text rendering needs
type Fonts struct {
	Face      text.Face
	TitleFace text.Face
}

// initFontFace sets up a font face from the embedded asset storage using specified parameters
func initFontFace(fontPath string, fontSize float64) (text.Face, error) {
	fontResource, err := embeddedAssets.Open(fontPath)
	if err != nil {
		return nil, err
	}

	fontSource, err := text.NewGoTextFaceSource(fontResource)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: fontSource,
		Size:   fontSize,
	}, nil
}

// SetupFontSystem initializes and returns the font collection needed for the UI
func SetupFontSystem() (*Fonts, error) {
	defaultFont, err := initFontFace(fontFaceRegular, 20)
	if err != nil {
		return nil, err
	}

	headerFont, err := initFontFace(fontFaceNarrow, 24)
	if err != nil {
		return nil, err
	}

	return &Fonts{
		Face:      defaultFont,
		TitleFace: headerFont,
	}, nil
}
