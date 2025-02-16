package ui

import (
	"image/color"
	"log"

	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	fontFaceRegular = "assets/Font/kenney-future.ttf"
	fontFaceNarrow  = "assets/Font/kenney-future-narrow.ttf"
)

var (
	gray = color.NRGBA{238, 244, 247, 255} // #eef4f7
)

type buttonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    text.Face
	padding widget.Insets
}

func loadFont(path string, size float64) (text.Face, error) {
	fontFile, err := embeddedAssets.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := text.NewGoTextFaceSource(fontFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}

func newImageFromFile(path string) (*ebiten.Image, error) {
	f, err := embeddedAssets.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	i, _, err := ebitenutil.NewImageFromReader(f)
	return i, err
}

func loadImageNineSlice(path string, centerWidth int, centerHeight int) (*eimage.NineSlice, error) {
	i, err := newImageFromFile(path)
	if err != nil {
		return nil, err
	}
	w := i.Bounds().Dx()
	h := i.Bounds().Dy()
	return eimage.NewNineSlice(i,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

func newButtonResources(fonts *fonts) (*buttonResources, error) {
	idle, err := loadImageNineSlice("assets/PNG/Extra/Default/button_rectangle_line.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("assets/PNG/Extra/Default/button_rectangle_depth_line.png", 12, 0)
	if err != nil {
		return nil, err
	}

	i := &widget.ButtonImage{
		Idle:    idle,
		Pressed: pressed,
	}

	return &buttonResources{
		image: i,

		text: &widget.ButtonTextColor{
			Idle:    color.Black,
			Pressed: color.Black,
		},

		face: fonts.face,

		padding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

type fonts struct {
	bigTitleFace text.Face
	face         text.Face
	titleFace    text.Face
	toolTipFace  text.Face
}

func loadFonts() (*fonts, error) {
	bigTitleFace, err := loadFont(fontFaceRegular, 24)
	if err != nil {
		return nil, err
	}

	fontFace, err := loadFont(fontFaceRegular, 20)
	if err != nil {
		return nil, err
	}

	titleFace, err := loadFont(fontFaceNarrow, 24)
	if err != nil {
		return nil, err
	}

	toolTipFace, err := loadFont(fontFaceNarrow, 20)
	if err != nil {
		return nil, err
	}

	return &fonts{
		bigTitleFace: bigTitleFace,
		face:         fontFace,
		titleFace:    titleFace,
		toolTipFace:  toolTipFace,
	}, nil
}
