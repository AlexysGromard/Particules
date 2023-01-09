package main

import (
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	// INTERACTION AVEC CONFIGURATION
	// Si on appuie sur la touche espace, on change de page
	if ebiten.IsKeyPressed(ebiten.KeySpace) && CurrentPage == configurationsPage {
		CurrentPage = particlesPage
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
			particles.Explosion(g.system.Content)
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
