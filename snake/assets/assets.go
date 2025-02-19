package assets

import (
	"embed"
	"io/fs"
	"log"

	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed *
var embeddedAssets embed.FS

// LoadEbitenImageFromAsset loads an image from the embedded assets and returns
// an ebiten.Image that can be used for rendering
func LoadEbitenImageFromAsset(assetPath string) (*ebiten.Image, error) {
	assetFile, err := embeddedAssets.Open(assetPath)
	if err != nil {
		return nil, err
	}
	defer assetFile.Close()

	ebitenImage, _, err := ebitenutil.NewImageFromReader(assetFile)

	return ebitenImage, err
}

// CreateScalableNineSliceImage creates a nine-slice image from a source image file
// with specified center dimensions for proper scaling
func CreateScalableNineSliceImage(assetPath string, sliceCenterWidth int, sliceCenterHeight int) (*image.NineSlice, error) {
	sourceImage, err := LoadEbitenImageFromAsset(assetPath)
	if err != nil {
		return nil, err
	}

	imageWidth := sourceImage.Bounds().Dx()
	imageHeight := sourceImage.Bounds().Dy()

	return image.NewNineSlice(sourceImage,
			[3]int{(imageWidth - sliceCenterWidth) / 2, sliceCenterWidth, imageWidth - (imageWidth-sliceCenterWidth)/2 - sliceCenterWidth},
			[3]int{(imageHeight - sliceCenterHeight) / 2, sliceCenterHeight, imageHeight - (imageHeight-sliceCenterHeight)/2 - sliceCenterHeight}),
		nil
}

// loadFontAsset returns the
func loadFontAsset(assetPath string) (fs.File, error) {
	assetFile, err := embeddedAssets.Open(assetPath)
	if err != nil {
		return nil, err
	}

	return assetFile, nil
}

// InitFontFace sets up a font face from the embedded asset storage using specified parameters
func InitFontFace(fontPath string, fontSize float64) (text.Face, error) {
	fontResource, err := loadFontAsset(fontPath)
	if err != nil {
		return nil, err
	}
	defer fontResource.Close()

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
