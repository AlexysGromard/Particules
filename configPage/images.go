package configPage

import (
	"image"
	_ "image/png"
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
func LoadAllImages(filePath string) error {
	// Get folder files
	files, err := os.ReadDir(filePath)
	if err != nil {
		return err
	}
	for _, f := range files {
		// Get image from file
		image, err := getImageFromFilePath(filePath + f.Name())
		if err != nil {
			return err
		}
		// add image to slice
		ImageList = append(ImageList, Images{name: f.Name(), image: image})
	}
	return nil
}

// GetImageFromName retourne une image à partir de son nom
// Elle ouvre l'image avec le module OS, elle load les images avec image.Decode et convertit pour Ebiten
// Entrée: nom de l'image
// Sortie : image *ebiten.Image
func getImageFromFilePath(filePath string) (img *ebiten.Image, err error) {
	// Open file with OS
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	// Decode image
	imgDecode, _, _ := image.Decode(file)
	// Convert for Ebiten
	image := ebiten.NewImageFromImage(imgDecode)
	if err != nil {
		return nil, err
	}
	return image, err
}

// FidnImage cherche une image grâce à son nom dans la liste des images
func findImage(images []Images, filename string) *ebiten.Image {
	// Boucle dans les images avec leurs noms
	for _, img := range images {
		// Trouve l'image avec le nom correspondant
		if img.name == filename {
			// Return img
			return img.image
		}
	}
	return nil
}
