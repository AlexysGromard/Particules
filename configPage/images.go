package configPage

import (
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Create struct for images
type Images struct {
	name  string
	image *ebiten.Image
}

// Create slice of images
var ImageList = []Images{}

// LoadImages charge en mémoire toutes les images du dossier assets/graphics
// Elle ouvre le dossier avec le module OS. Elle boucle sur toutes les images et les ajoute à une slice
// Entrée: aucun
// Sortie: aucun
func LoadImages() {
	// Get folder files
	files, err := ioutil.ReadDir("./assets/graphics")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// Get image from file
		image, err := getImageFromFilePath("./assets/graphics/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		// add image to slice
		ImageList = append(ImageList, Images{name: f.Name(), image: image})
	}
}

// GetImageFromName retourne une image à partir de son nom
// Elle ouvre l'image avec le module OS, elle load les images avec image.Decode et convertit pour Ebiten
// Entrée: nom de l'image
// Sortie : image *ebiten.Image
func getImageFromFilePath(filePath string) (img *ebiten.Image, err error) {
	// Open file with OS
	file, err := os.Open(filePath)
	// Decode image
	imgDecode, _, _ := image.Decode(file)
	// Convert for Ebiten
	image := ebiten.NewImageFromImage(imgDecode)
	if err != nil {
		log.Fatal(err)
	}
	return image, err
}
