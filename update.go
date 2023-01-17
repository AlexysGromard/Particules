package main

import (
	"project-particles/config"
	"project-particles/configPage"
	//"project-particles/particles"
	"project-particles/particles/ParticleModification"
	
	"github.com/hajimehoshi/ebiten/v2"
)

var configPageScrollY int

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	// Scroll sur la page de configuration
	if CurrentPage == configurationsPage && config.General.WindowSizeY < 800 {
		_, y := ebiten.Wheel()
		hiddenElementsSizeY := 1050 - config.General.WindowSizeY
		if y > 0 && configPageScrollY < 0 || y < 0 && configPageScrollY > -hiddenElementsSizeY {
			configPageScrollY += int(y * 8)
			configPage.ScrollY = int(y * 8)
		} else {
			configPage.ScrollY = 0
		}
	} else if configPageScrollY != 0 {
		// Remettre les valeurs par défaut
		configPage.ScrollY -= configPageScrollY
		configPageScrollY = 0
	} else {
		configPage.ScrollY = 0
	}

	// INTERACTION AVEC CONFIGURATION
	// Si on appuie sur la touche espace, on change de page
	if ebiten.IsKeyPressed(ebiten.KeyEnter) && CurrentPage == configurationsPage {
		CurrentPage = particlesPage
		// Save la configuration
		configPage.SaveConfig()
	} else if ebiten.IsKeyPressed(ebiten.KeyEscape) && CurrentPage == particlesPage {
		CurrentPage = configurationsPage
	}
	if config.General.Interaction {
		// INTERACTION AVEC PARTICULES
		// Deplacement de la zone de spawn
		// Si fleche haut est appuyee, on diminue la coordonnee Y de la zone de spawn
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

	g.system.Update()

	return nil
}
