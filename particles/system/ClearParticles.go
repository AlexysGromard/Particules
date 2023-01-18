package system

func (s *System) ClearParticles() {
	NuméroPaticule := s.Content.Front()
	for NuméroPaticule != nil {
		ParticuleSuivante := NuméroPaticule.Next()
		s.KillParticule(NuméroPaticule)
		NuméroPaticule = ParticuleSuivante
	}
}
