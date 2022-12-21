# Particules
Particule est un projet créé en Go (golang) pour le cours d'initiation au développement à <a href="https://iutnantes.univ-nantes.fr">l'IUT de Nantes</a>.

## Pour commencer
### Pré-requis
- Go <a href="https://go.dev/doc/install">(Lien vers l'instalation de Go)</a>
- Git <a href="https://git-scm.com/book/fr/v2/Démarrage-rapide-Installation-de-Git">(Lien vers l'instalation de Git)</a>

### Installation
1. Cloner le projet
```bash
git clone https://gitlab.univ-nantes.fr/particules/particules.git
```
2. Se déplacer dans le dossier du projet
```bash
cd particules
```
3. Compiler le projet
```bash
go build
```
4. Lancer le projet
```bash
./particules
```
(ou ```particules.exe``` sur Windows)

## Configuration
Dans ce projet, il se trouve un fichier intitulé ```config.json``` qui permet de configurer le projet.

### Configuration de la fenêtre
```json
	"WindowTitle": "Project particles",
	"WindowSizeX": 800,
	"WindowSizeY": 600,
	"ParticleImage": "assets/particle.png",
````
- ```WindowTitle``` : Titre de la fenêtre
- ```WindowSizeX``` : Largeur de la fenêtre
- ```WindowSizeY``` : Hauteur de la fenêtre
- ```ParticleImage``` : Chemin vers l'image de la particule
### Debug
```json
	"Debug": true,
````
- ```Debug``` : Affiche ou non les informations de debug
### Génération des particules
```json
	"InitNumParticles": 100,
	"RandomSpawn": false,
	"SpawnX": 400,
	"SpawnY": 300,
	"SpawnRate": 2,
````
- ```InitNumParticles``` : Nombre de particules à générer au lancement du programme
- ```RandomSpawn``` : Si ```true```, les particules sont générées aléatoirement sur la fenêtre, sinon elles sont générées à la position ```SpawnX``` et ```SpawnY```
- ```SpawnX``` : Position X de la génération des particules
- ```SpawnY``` : Position Y de la génération des particules
- ```SpawnRate``` : Nombre de particules générées par seconde
### Propriétés des particules
```json
	"ScaleX": 1,
	"ScaleY": 1,
	"Opacity": 0.75,
	"ColorRed": 0.5,
	"ColorGreen": 0.5,
	"ColorBlue": 1,
````
- ```ScaleX``` : Taille de la particule en X
- ```ScaleY``` : Taille de la particule en Y
- ```Opacity``` : Opacité de la particule
- ```ColorRed``` : Couleur de la particule en rouge (1 = 255 en RGB)
- ```ColorGreen``` : Couleur de la particule en vert (1 = 255 en RGB)
- ```ColorBlue``` : Couleur de la particule en bleu (1 = 255 en RGB)
### Comportement de la particule 
```json
	"SpeedType" : 2,
	"KillParticlesOutside" : true,
	"Gravity" : 1,
	"MarginOutsideScreen": 1,
	"HaveLife": true,
	"RandomLife": true,
	"Life": 600
````
- ```SpeedType``` : Type de vitesse de la particule
    - 1 : Vitesse faible
    - 2 : Vitesse moyenne
    - 3 : Vitesse forte
- ```KillParticlesOutside``` : Si ```true```, les particules sont supprimées si elles sortent de l'écran
- ```Gravity``` : Force de la gravité
- ```MarginOutsideScreen``` : Marge de la fenêtre pour la suppression des particules
- ```HaveLife``` : Si ```true```, les particules ont une durée de vie
- ```RandomLife``` : Si ```true```, la durée de vie des particules est aléatoire
- ```Life``` : Durée de vie des particules