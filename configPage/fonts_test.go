package configPage

import "testing"

// la fonction loadFontTitle() test si le fichier de la police de caractères existe
// et si il est bien chargé
// Elle retourne une erreur si le fichier n'existe pas ou si il n'est pas chargé
func TestLoadFontRegular(t *testing.T) {
	err := loadFontRegular()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

// la fonction loadFontBold() test si le fichier de la police de caractères existe
// et si il est bien chargé
// Elle retourne une erreur si le fichier n'existe pas ou si il n'est pas chargé
func TestLoadFontBold(t *testing.T) {
	err := loadFontBold()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

// la fonction loadFontTitle() test si le fichier de la police de caractères existe
// et si il est bien chargé
// Elle retourne une erreur si le fichier n'existe pas ou si il n'est pas chargé
func TestLoadFontTitle(t *testing.T) {
	err := loadFontTitle()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
