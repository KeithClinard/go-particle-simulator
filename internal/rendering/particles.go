package rendering

import (
	"log"

	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var particleSprite *ebiten.Image
var planetDrawOptions = &ebiten.DrawImageOptions{}
var particleDrawOptions = &ebiten.DrawImageOptions{}

func init() {
	sprite, _, err := ebitenutil.NewImageFromFile("assets/particle.png")
	particleSprite = sprite
	if err != nil {
		log.Fatal(err)
	}
	planetDrawOptions.ColorM.Scale(1, 0, 0, 1)

	particleDrawOptions.GeoM.Scale(.04, .04)
}

func DrawParticles(gameState *models.GameState, screen *ebiten.Image) {
	for _, particle := range gameState.Particles {
		drawParticle(particle, screen)
	}
	for _, particle := range gameState.Planets {
		drawPlanet(particle, screen)
	}
}

func drawPlanet(particle *models.Particle, screen *ebiten.Image) {
	planetDrawOptions.GeoM.Reset()
	particleRadius := 0 - particle.Size/2.0
	planetDrawOptions.GeoM.Translate(particle.Position.X+particleRadius, particle.Position.Y+particleRadius)
	screen.DrawImage(particleSprite, planetDrawOptions)
}

func drawParticle(particle *models.Particle, screen *ebiten.Image) {
	particleDrawOptions.GeoM.Reset()
	particleDrawOptions.GeoM.Scale(.04, .04)
	particleDrawOptions.GeoM.Translate(particle.Position.X-2, particle.Position.Y-2)
	screen.DrawImage(particleSprite, particleDrawOptions)
}
