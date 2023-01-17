package configPage

import (
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// Importer les deux polices de caractères

const (
	// Chemin des fonts
	fontFaceRegular = "assets/fonts/Roboto-Regular.ttf"
	fontFaceBold    = "assets/fonts/Roboto-Bold.ttf"
)

var (
	// RobotoRegular
	RobotoRegularFont  *truetype.Font
	RobotoRegularFontF font.Face
	// RobotoBold
	RobotoBoldFont  *truetype.Font
	RobotoBoldFontF font.Face
	// PageTitle
	PageTitle  *Text
	PageTitleF font.Face
)

// loadFontRegular charge une police de caractères regular
func loadFontRegular() error {
	b, err := os.ReadFile(fontFaceRegular)
	if err != nil {
		return err
	}
	RobotoRegularFont, err = truetype.Parse(b)
	if err != nil {
		return err
	}
	RobotoRegularFontF = truetype.NewFace(RobotoRegularFont, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return nil
}

// LoadFontBold charge une police de caractères bold
func loadFontBold() error {
	b, err := os.ReadFile(fontFaceBold)
	if err != nil {
		return err
	}
	RobotoBoldFont, err = truetype.Parse(b)
	if err != nil {
		return err
	}
	RobotoBoldFontF = truetype.NewFace(RobotoBoldFont, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return nil
}

// LoadFontTitle
func loadFontTitle() error {
	b, err := os.ReadFile(fontFaceBold)
	if err != nil {
		return err
	}
	RobotoBoldFont, err = truetype.Parse(b)
	if err != nil {
		return err
	}
	PageTitleF = truetype.NewFace(RobotoBoldFont, &truetype.Options{
		Size:    25,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return nil
}
