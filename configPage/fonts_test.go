package configPage

import "testing"

// la fonction loadFontTitle() test si le fichier de la police de caractères existe
// et si il est bien chargé
// Elle retourne une erreur si le fichier n'existe pas ou si il n'est pas chargé
func TestLoadFontRegular(t *testing.T) {
	err := loadFontRegular("../assets/fonts/Roboto-Regular.ttf")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

// la fonction loadFontBold() test si le fichier de la police de caractères existe
// et si il est bien chargé
// Elle retourne une erreur si le fichier n'existe pas ou si il n'est pas chargé
func TestLoadFontBold(t *testing.T) {
	err := loadFontBold("../assets/fonts/Roboto-Bold.ttf")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

// la fonction loadFontTitle() test si le fichier de la police de caractères existe
// et si il est bien chargé
// Elle retourne une erreur si le fichier n'existe pas ou si il n'est pas chargé
func TestLoadFontTitle(t *testing.T) {
	err := loadFontTitle("../assets/fonts/Roboto-Bold.ttf")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestLoadFontsFail(t *testing.T) {
	err1 := loadFontRegular("toto/tata")
	err2 := loadFontBold("toto/tata")
	err3 := loadFontTitle("blablabla/toto")
	if err1 == nil || err2 == nil || err3 == nil {
		t.Errorf("Error: Le chargement d'une police qui n'existe pas retourne une erreur")
	}
}
