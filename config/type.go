package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawn              bool
	SpawnX, SpawnY           int
	SpawnRate				 float64
	// Particule properties
	ScaleX, ScaleY           float64
	Opacity				  	 float64
	ColorRed, ColorGreen, ColorBlue float64
	SpeedType				 int
	KillParticlesOutside	 bool
	Gravity					 float64
	MarginOutsideScreen		 float64
}

var General Config
