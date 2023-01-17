package configPage

import (
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// Variables globales des polices de caractères
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
// Taille de la police: 20
// Type de police: regular
func loadFontRegular(path string) error {
	// Charge le fichier de la police de caractères .ttf
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	// Analyse et charge la police de caractères dans la variable
	RobotoRegularFont, err = truetype.Parse(file)
	if err != nil {
		return err
	}
	// Création de la police de caractères
	RobotoRegularFontF = truetype.NewFace(RobotoRegularFont, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return nil
}

// LoadFontBold charge une police de caractères bold
// Taille de la police: 20
// Type de police: bold
func loadFontBold(path string) error {
	// Charge le fichier de la police de caractères .ttf
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	// Analyse et charge la police de caractères dans la variable
	RobotoBoldFont, err = truetype.Parse(file)
	if err != nil {
		return err
	}
	// Création de la police de caractères
	RobotoBoldFontF = truetype.NewFace(RobotoBoldFont, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return nil
}

// LoadFontTitle charge une police de caractères bold
// Taille de la police: 25
// Type de police: bold
func loadFontTitle(path string) error {
	// Charge le fichier de la police de caractères .ttf
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	// Analyse et charge la police de caractères dans la variable
	RobotoBoldFont, err = truetype.Parse(file)
	if err != nil {
		return err
	}
	// Création de la police de caractères
	PageTitleF = truetype.NewFace(RobotoBoldFont, &truetype.Options{
		Size:    25,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return nil
}
