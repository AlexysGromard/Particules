package system

import (
	"project-particles/config"
	"testing"
)

// Le test TestNewSystem vérifie que la fonction NewSystem initialise bien le system avec le nombre de particule
// défini dans la configuration. Ici 123 particules
// Si ce n'est pas le cas, le test échoue
func TestNewSystem(t *testing.T) {

	config.General.InitNumParticles = 123

	system := NewSystem()

	if system.Content.Len() != 123 {
		t.Error("il devrait y avoir 123 particule mais il y en à ", system.Content.Len())
	}
}
