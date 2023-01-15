package configPage

import (
	"encoding/json"
	"os"
	"project-particles/config"
	"strings"
)

func SaveConfig() {
	// Récupérer General de type Config
	general := config.General
	// Ouvrir le fichier de config en écriture écrasant le contenu
	file, err := os.OpenFile("config.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	// écrire dans le fichier de la même manière que pour la lecture
	jsonContent, _ := json.Marshal(general)
	jsonContentString := string(jsonContent)
	jsonContentString = strings.Replace(jsonContentString, ",", ",\n", -1)
	jsonContentString = strings.Replace(jsonContentString, "{", "{\n", -1)
	jsonContentString = strings.Replace(jsonContentString, "}", "\n}", -1)
	file.WriteString(jsonContentString)

	// Fermeture du fichier
	file.Close()
}
