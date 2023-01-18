package system

// La petite fonction ClearParticles permet de supprimer toutes les particules du système
// Elle est appelée par le bouton Clear sur la page de configuration
func (s *System) ClearParticles() {
	NuméroPaticule := s.Content.Front()
	for NuméroPaticule != nil {
		ParticuleSuivante := NuméroPaticule.Next()
		s.killParticule(NuméroPaticule)
		NuméroPaticule = ParticuleSuivante
	}
}
