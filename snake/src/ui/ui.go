package ui

import (
	"bytes"
	"image/color"
	"log"

	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	color1 = color.NRGBA{239, 121, 138, 255} // #ef798a
	color2 = color.NRGBA{247, 169, 168, 255} // #f7a9a8
	color3 = color.NRGBA{97, 63, 117, 255}   // #613f75
	color4 = color.NRGBA{229, 195, 209, 255} // #e5c3d1
)

func loadFont(size float64) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}

func loadButtonImage() (*widget.ButtonImage, error) {
	transparent := eimage.NewNineSliceColor(color.Transparent)

	return &widget.ButtonImage{
		Idle:    transparent,
		Hover:   transparent,
		Pressed: transparent,
	}, nil
}
