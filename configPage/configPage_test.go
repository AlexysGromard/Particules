package configPage

import (
	"testing"
)

// la fonction loadImages() sert à charger les images des différents composants
// de la page de configuration
// Cette fonction utilise le slice ImageList qui contient toutes les images
// Ce test vérifie que les images ont bien été chargées
func TestLoadImages(t *testing.T) {
	// LoadAllImages charge toutes les images dans le slice ImageList
	LoadAllImages("../assets/graphics/")
	// Appel de la fonction à tester
	loadImages()
	if checkboxImages == nil {
		t.Error("Les images des checkbox n'ont pas été chargées")
	} else if buttonImages == nil {
		t.Error("Les images des bouttons n'ont pas été chargées")
	} else if sliderImages == nil {
		t.Error("Les images des sliders n'ont pas été chargées")
	} else if numberInputImages == nil {
		t.Error("Les images des radios n'ont pas été chargées")
	}
}

// TestCheckImagesCheckbox vérifie que le nombre d'images des checkbox est correct
// Le nombre d'images des checkbox doit être 6
func TestCheckImagesCheckbox(t *testing.T) {
	// LoadAllImages charge toutes les images dans le slice ImageList
	LoadAllImages("../assets/graphics/")
	// Appel de la fonction à tester
	loadImages()
	if len(checkboxImages) != 6 {
		t.Error("Le nombre d'images des checkbox est incorrect")
	}
}

// La test TestCreatItems vérifie que les composants de la page de configuration
// ont bien été créés.
// Elle vérifie en testant une dizaine de composants
// Si un composant == nil, alors il n'est pas créé et le test échoue
func TestCreatItems(t *testing.T) {
	createItems()
	if speedType == nil {
		t.Error("Le type de vitesse n'a pas été créé")
	} else if opacityText == nil {
		t.Error("Le texte d'opacité n'a pas été créé")
	} else if spawnRateValue == nil {
		t.Error("La valeur du taux de spawn n'a pas été créé")
	} else if leaveGamebutton == nil {
		t.Error("Le boutton de quitter le jeu n'a pas été créé")
	} else if minColorBlue == nil {
		t.Error("Le bleu minimum n'a pas été créé")
	} else if maxColorBlue == nil {
		t.Error("Le bleu maximum n'a pas été créé")
	} else if minColorGreen == nil {
		t.Error("Le vert minimum n'a pas été créé")
	} else if maxColorGreen == nil {
		t.Error("Le vert maximum n'a pas été créé")
	} else if minColorRed == nil {
		t.Error("Le rouge minimum n'a pas été créé")
	} else if maxColorRed == nil {
		t.Error("Le rouge maximum n'a pas été créé")
	} else if PlayButton == nil {
		t.Error("Le boutton de lancer le jeu n'a pas été créé")
	}
}

// Le test TestLoadFont vérifie que les polices de caractères ont bien été chargées
// Si la fonction loadFont() retourne une erreur, alors le test échoue
func TestLoadFont(t *testing.T) {
	// Appel de la fonction à tester
	err := loadFont("../assets/fonts/")
	if err != nil {
		t.Error("Les polices de caractères n'ont pas été chargées alors qu'elles devraient l'être")
	}
}

// Le test TestLoadFontError vérifie que la fonction loadFont() retourne une erreur
// si le chemin vers les polices de caractères est faux
// Elle met en paramètre un chemin faux
func TestLoadFontError(t *testing.T) {
	// Appel de la fonction à tester
	err := loadFont("../pas/le/bon/chemin/")
	if err == nil {
		t.Error("Les polices de caractères ont été chargées alors que le chemin est faux")
	}
}
