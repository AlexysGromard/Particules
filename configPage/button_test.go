package configPage

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// TestCreateBtn test la création d'un boutton
// Cette fonction charge les images et la police de caractère
// puis elle crée un boutton et vérifie qu'il n'est pas nil
// Si le boutton est nil, alors le test échoue
func TestCreateBtn(t *testing.T) {
	// Load font
	_ = loadFontRegular("./assets/fonts/Roboto-Regular.ttf")
	// Load images
	buttonImages = []*ebiten.Image{
		findImage(ImageList, "button-idle.png"),    // Default
		findImage(ImageList, "button-pressed.png"), // Pressed
		findImage(ImageList, "button-hover.png")}   // Hover

	// Création d'un boutton
	btn := newButton(50, 100, 170, 50, buttonImages, "test", RobotoRegularFontF, func() {})
	if btn == nil {
		t.Error("Le boutton n'a pas été créé")
	}
}
