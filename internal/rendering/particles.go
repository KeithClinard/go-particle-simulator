package rendering

import (
	"log"

	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var particleSprite *ebiten.Image
var particleDrawOptions = &ebiten.DrawImageOptions{}

func init() {
	sprite, _, err := ebitenutil.NewImageFromFile("assets/particle.png")
	particleSprite = sprite
	if err != nil {
		log.Fatal(err)
	}
}

func DrawParticles(gameState *models.GameState, screen *ebiten.Image) {
	for _, particle := range gameState.Particles {
		particleDrawOptions.GeoM.Reset()
		particleScaleFactor := particle.Size / 100
		particleDrawOptions.GeoM.Scale(particleScaleFactor, particleScaleFactor)
		particleRadius := 0 - particle.Size/2
		particleDrawOptions.GeoM.Translate(particleRadius, particleRadius)
		particleDrawOptions.GeoM.Translate(particle.Position.X, particle.Position.Y)
		screen.DrawImage(particleSprite, particleDrawOptions)
	}
}