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
	if config.General.Interaction {
		// Deplacement de la zone de spawn
		// Si fleche haut est appuyee, on diminue la coordonnee Y de la zone de spawn
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			config.General.SpawnY -= 3
		}
		// Si fleche bas est appuyee, on augmente la coordonnee Y de la zone de spawn
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			config.General.SpawnY += 3
		}
		// Si fleche gauche est appuyee, on diminue la coordonnee X de la zone de spawn
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			config.General.SpawnX -= 3
		}
		// Si fleche droite est appuyee, on augmente la coordonnee X de la zone de spawn
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			config.General.SpawnX += 3
		}

		// Explosion
		// Si espace est appuyee, on appelle la fonction Explosion du systeme de particules
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			particles.Explosion(g.system.Content)
		}
	}
	if config.General.FollowMouse {
		// On recupere les coordonnees de la souris et on met à jour la zone de spawn
		config.General.SpawnX, config.General.SpawnY = ebiten.CursorPosition()
	}

	g.system.Update()

	return nil
}
