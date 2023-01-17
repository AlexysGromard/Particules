package configPage

import (
	"testing"
)

// TestLoadImages vérifie que la fonction LoadImages ne retourne pas d'erreur
// Elle envoie un chemin d'accès à une image existante
// La fonction LoadImages ne doit pas retourner d'erreur
func TestGetImageFromFilePath(t *testing.T) {
	_, err := getImageFromFilePath("../assets/graphics/checkbox-disabled.png")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

// TestLoadAllImages vérifie que la fonction LoadAllImages ne retourne pas d'erreur
// Elle envoie un chemin d'accès à un dossier contenant des images existantes
// La fonction LoadAllImages ne doit pas retourner d'erreur
func TestLoadAllImages(t *testing.T) {
	err := LoadAllImages("../assets/graphics/")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

// TestFindImage vérifie que la fonction findImage retourne nil
// Elle envoie un nom d'image qui n'existe pas
func TestFindImageNotExist(t *testing.T) {
	// LoadAllImages charge toutes les images dans ImageList (slice dans images.go)
	_ = LoadAllImages("../assets/graphics/")
	image := findImage(ImageList, "pas-une-image.pdf")
	if image != nil {
		t.Errorf("Error: pas-une-image.pdf n'existe pas.")
	}
}

// TestFindImage vérifie que la fonction findImage retourne une image
// Elle envoie un nom d'image qui existe
// La fonction findImage doit retourner une image
func TestFindImage(t *testing.T) {
	// LoadAllImages charge toutes les images dans ImageList (slice dans images.go)
	_ = LoadAllImages("../assets/graphics/")
	image := findImage(ImageList, "checkbox-disabled.png")
	if image == nil {
		t.Errorf("Error: checkbox-disabled.png existe.")
	}
}

// TestFindAllImages vérifie que la fonction findImage retourne une image
// Elle envoie tous les noms d'images qui existent et qui sont dans ImageList
// Elle ne doit pas retourner d'erreur
func TestFindAllImages(t *testing.T) {
	// LoadAllImages charge toutes les images dans ImageList (slice dans images.go)
	_ = LoadAllImages("../assets/graphics/")
	for _, image := range ImageList {
		if findImage(ImageList, image.name) == nil {
			t.Errorf("Error: %v existe.", image.name)
		}
	}

}

// TestGetImageFromFilePathNotExist vérifie que la fonction getImageFromFilePath retourne une erreur
// Elle envoie un chemin d'accès à une image qui n'existe pas
// La fonction getImageFromFilePath doit retourner une erreur "no such file or directory"
func TestGetImageFromFilePathNotExist(t *testing.T) {
	_, err := getImageFromFilePath("../assets/graphics/pas-image.pdf")
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}
