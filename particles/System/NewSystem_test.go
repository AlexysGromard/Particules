package system

import (
	"project-particles/config"
	"testing"
)

func TestNewSystem(t *testing.T) {

	config.General.InitNumParticles = 123

	system := NewSystem()

	if system.Content.Len() != 123 {
		t.Error("il devrait y avoir 123 particule mais il y en à ", system.Content.Len())
	}
}
