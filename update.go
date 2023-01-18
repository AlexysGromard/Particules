package main

import (
	"project-particles/config"
	"project-particles/configPage"

	//"project-particles/particles"
	"project-particles/particles/ParticleModification"

	"github.com/hajimehoshi/ebiten/v2"
)

var configPageScrollY int // Variable qui contient le scroll de la page de configuration

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	// SCROLL SUR LA PAGE DE CONFIGURATION
	// Calculer la taille des éléments cachés
	hiddenElementsSizeY := 1080 - config.General.WindowSizeY
	// Si on est sur la page de configuration et qu'il y a des éléments cachés, on scroll
	if CurrentPage == configurationsPage && hiddenElementsSizeY > 0 {
		_, y := ebiten.Wheel() // Valeur du scroll
		if (y < 0 && -(configPageScrollY+int(y*8)) < hiddenElementsSizeY) || (y > 0 && -(configPageScrollY+int(y*8)) > 0) {
			// Si on scroll vers le haut ou vers le bas et que le scroll ne dépasse pas les limites
			configPageScrollY += int(y * 8)
			configPage.ScrollY = int(y * 8)
		} else if y < 0 && -(configPageScrollY+int(y*8)) > hiddenElementsSizeY && -configPageScrollY < hiddenElementsSizeY {
			// Si on scroll vers le haut et que le scroll dépasse les limites
			configPage.ScrollY = -hiddenElementsSizeY - configPageScrollY
			configPageScrollY = -hiddenElementsSizeY
		} else if y > 0 && -(configPageScrollY+int(y*8)) < 0 && -configPageScrollY > 0 {
			// Si on scroll vers le bas et que le scroll dépasse les limites
			configPage.ScrollY = -configPageScrollY
			configPageScrollY = 0
		} else {
			configPage.ScrollY = 0
		}
	} else if configPageScrollY != 0 {
		// Remettre les valeurs par défaut car il n'y a plus de scroll
		configPage.ScrollY -= configPageScrollY
		configPageScrollY = 0
	}

	// INTERACTION AVEC CONFIGURATION
	// Si on appuie sur la Enter ou le bouton playButton, on va sur la page de particules
	// Si on appuie sur la touche Echap, on va sur la page de configuration
	if configPage.PlayButton != nil {
		if (ebiten.IsKeyPressed(ebiten.KeyEnter) || configPage.PlayButton.Pressed) && CurrentPage == configurationsPage {
			CurrentPage = particlesPage
		} else if ebiten.IsKeyPressed(ebiten.KeyEscape) && CurrentPage == particlesPage {
			CurrentPage = configurationsPage
		}
	}
	if config.General.Interaction {
		// INTERACTION AVEC PARTICULES
		// Deplacement de la zone de spawn
		// Si fleche haut est appuyee, on diminue la coordonnee Y de la zone de spawn
		if !config.General.SpawnCenter {
			if ebiten.IsKeyPressed(ebiten.KeyUp) && config.General.SpawnY > 0 {
				config.General.SpawnY -= 3
			}
			// Si fleche bas est appuyee, on augmente la coordonnee Y de la zone de spawn
			if ebiten.IsKeyPressed(ebiten.KeyDown) && config.General.SpawnY < config.General.WindowSizeY {
				config.General.SpawnY += 3
			}
			// Si fleche gauche est appuyee, on diminue la coordonnee X de la zone de spawn
			if ebiten.IsKeyPressed(ebiten.KeyLeft) && config.General.SpawnX > 0 {
				config.General.SpawnX -= 3
			}
			// Si fleche droite est appuyee, on augmente la coordonnee X de la zone de spawn
			if ebiten.IsKeyPressed(ebiten.KeyRight) && config.General.SpawnX < config.General.WindowSizeX {
				config.General.SpawnX += 3
			}
		}
		// Explosion
		// Si espace est appuyee, on appelle la fonction Explosion du systeme de particules
		if ebiten.IsKeyPressed(ebiten.KeySpace) && CurrentPage == particlesPage {
			ParticleModification.Explosion(g.system.Content)
		}
		// Tourbillon
		// Si T est appuyee, on appelle la fonction MakeWhirlwind du systeme de particules
		if ebiten.IsKeyPressed(ebiten.KeyT) {
			ParticleModification.MakeWhirlwind(g.system.Content)
		}
	}
	if config.General.FollowMouse {
		// On recupere les coordonnees de la souris
		x, y := ebiten.CursorPosition()
		// Si les coordonnees sont dans la fenetre, on met a jour les coordonnees de la zone de spawn
		if x >= 0 && x <= config.General.WindowSizeX && y >= 0 && y <= config.General.WindowSizeY {
			// On met a jour les coordonnees de la zone de spawn
			config.General.SpawnX = x
			config.General.SpawnY = y
		}
	}

	if CurrentPage == particlesPage {
		// Appel de la fonction Update du systeme de particules
		g.system.Update()
	}

	return nil
}
