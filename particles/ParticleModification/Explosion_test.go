package ParticleModification

import (
	"container/list"
	"math"
	"project-particles/particles"
	"testing"
)

// Le test TestExplosionNombreParticuleCree vérifie que la fonction Explosion crée bien 100 particules
func TestExplosionNombreParticuleCree(t *testing.T) {
	l := list.New()

	Explosion(l)

	if l.Len() == 0 {
		t.Error("aucune particule n'a été crée")
	} else if l.Len() != 100 {
		t.Error("il n'y a pas eu 100 particules qui ont été crées")
	}
}

// Le test TestExplosionVitesseConforme vérifie que la fonction Explosion crée bien des particules avec une vitesse entre 3 et 5
// Si une particule n'a pas une vitesse entre 3 et 5, le test échoue
// Si un element n'est pas une particule, le test échoue
func TestExplosionVitesseConforme(t *testing.T) {
	l := list.New()

	Explosion(l)

	for NuméroParticule := l.Front(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Next() {
		particule, ok := NuméroParticule.Value.(*particles.Particle)

		if ok {
			v_Réelle := math.Sqrt(particule.SpeedX*particule.SpeedX + particule.SpeedY*particule.SpeedY) //la fonction Sqrt du math accepte considère que Sqrt(0) == 0

			if v_Réelle < float64(3) || v_Réelle > float64(5) {
				t.Error("Une particule n'a pas une vitesse entre 3 et 5")
			}
		} else {
			t.Error("Un element de la liste n'est pas une particule")
		}
	}
}
