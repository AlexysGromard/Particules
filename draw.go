package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok {
			if !((p.PositionX < 0-config.General.MarginOutsideScreen || p.PositionX > float64(config.General.WindowSizeX) + config.General.MarginOutsideScreen) || (p.PositionY < 0-config.General.MarginOutsideScreen || p.PositionY > float64(config.General.WindowSizeY) + config.General.MarginOutsideScreen)){
				options := ebiten.DrawImageOptions{}
				options.GeoM.Rotate(p.Rotation)
				options.GeoM.Scale(p.ScaleX, p.ScaleY)
				options.GeoM.Translate(p.PositionX, p.PositionY)
				options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
				screen.DrawImage(assets.ParticleImage, &options)
			}
		}
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
	}

}
