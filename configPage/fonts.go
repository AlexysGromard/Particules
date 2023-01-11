package configPage

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// Importer les deux polices de caractères
const (
	fontFaceRegular = "assets/fonts/Roboto-Regular.ttf"
	fontFaceBold    = "assets/fonts/Roboto-Bold.ttf"
)

// Créer une structure pour les polices
type fonts struct {
	face         font.Face
	titleFace    font.Face
	bigTitleFace font.Face
	toolTipFace  font.Face
}

// LoadFonts charge les polices de caractères
// Entrée: aucun
// Sortie: fonts, error
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

// LoadFont charge une police de caractères à partir d'un fichier et d'une taille
// Entrée: chemin du fichier, taille de la police
// Sortie: font.Face, error
func loadFont(path string, size float64) (font.Face, error) {
	// Lecture du fichier
	fontData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Copie la police dans une variable
	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}
	// Retourne la police
	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
