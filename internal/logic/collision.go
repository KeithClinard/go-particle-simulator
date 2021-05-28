package logic

import (
	"math"

	"github.com/KeithClinard/go-particle-simulator/internal/models"
)

func HandleCollision(gameState *models.GameState) {
	numParticles := len(gameState.Particles)
	for i := 0; i < numParticles-1; i++ {
		particle1 := gameState.Particles[i]
		for j := i + 1; j < numParticles; j++ {
			particle2 := gameState.Particles[j]
			if particle1.ShouldBeDestroyed || particle2.ShouldBeDestroyed {
				continue
			}
			maxRadius := math.Max(particle1.Size/2.0, particle1.Size/2.0)
			displacement := particle2.Position.Clone().Subtract(*particle1.Position).Length()

			if displacement <= maxRadius {
				CollideParticles(particle1, particle2)
			}
		}
	}
}

func CollideParticles(p1 *models.Particle, p2 *models.Particle) {
	if p1.Size > p2.Size {
		p1.Absorb(p2)
	} else {
		p2.Absorb(p1)
	}
}
