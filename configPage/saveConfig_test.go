package configPage

import (
	"bufio"
	"os"
	"project-particles/config"
	"testing"
)

// La fonction TestSaveConfig permet de tester la fonction SaveConfig
// Elle vérifie que la fonction ne retourne pas d'erreur
// Elle modifie la valeur de Debug à true
func TestSaveConfig(t *testing.T) {
	// Mettre Debug à true
	config.General.Debug = true
	err := SaveConfig("../config.json")
	if err != nil {
		t.Error(err)
	}
}

// La fonction TestValueInJSON permet de tester la fonction SaveConfig
// Elle vérifie que le fichier config.json contient bien la valeur Debug à false
func TestValueInJSON(t *testing.T) {
	// Mettre Debug à true
	config.General.Debug = false
	_ = SaveConfig("../config.json")
	file, err := os.Open("../config.json")
	if err != nil {
		t.Error(err)
	}
	// Lit le fichier jusqu'à "Debug":false,
	// Si la valeur n'est pas présente, le test échoue
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "\"Debug\":false," {
			return
		}
	}
	t.Error("La valeur Debug n'est pas présente dans le fichier config.json")
}

// La fonction TestError permet de tester la fonction SaveConfig
// Elle vérifie que la fonction retourne bien une erreur si le chemin est vide
func TestError(t *testing.T) {
	err := SaveConfig("")
	if err == nil {
		t.Error("La fonction SaveConfig ne retourne pas d'erreur alors qu'elle devrait.")
	}
}
