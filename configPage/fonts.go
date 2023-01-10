package configPage

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	fontFaceRegular = "assets/fonts/Roboto-Regular.ttf"
	fontFaceBold    = "assets/fonts/Roboto-Bold.ttf"
)

type fonts struct {
	face         font.Face
	titleFace    font.Face
	bigTitleFace font.Face
	toolTipFace  font.Face
}

func loadFonts() (*fonts, error) {
	// Font face 20
	fontFace, err := loadFont(fontFaceRegular, 20)
	if err != nil {
		return nil, err
	}

	// Title font face
	titleFontFace, err := loadFont(fontFaceBold, 24)
	if err != nil {
		return nil, err
	}

	// Big title font face
	bigTitleFontFace, err := loadFont(fontFaceBold, 28)
	if err != nil {
		return nil, err
	}

	// Tool tip font face
	toolTipFace, err := loadFont(fontFaceRegular, 15)
	if err != nil {
		return nil, err
	}
	// Return all fonts
	return &fonts{
		face:         fontFace,
		titleFace:    titleFontFace,
		bigTitleFace: bigTitleFontFace,
		toolTipFace:  toolTipFace,
	}, nil
}

func loadFont(path string, size float64) (font.Face, error) {
	fontData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
