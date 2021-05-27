package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
)

var gravitationalConstant = 1000.0
var maxAllowedAcceleration = 100.0

func MoveAllParticles(gameState *models.GameState) {
	for _, particle := range gameState.Particles {
		particle.Move()
	}
}

func ApplyParticleGravity(gameState *models.GameState) {
	for _, particle1 := range gameState.Particles {
		accelerationSum := &models.Vector{
			X: 0,
			Y: 0,
		}
		for _, particle2 := range gameState.Particles {
			if particle1 == particle2 {
				continue
			}
			displacement := particle2.Position.Clone().Subtract(*particle1.Position)
			displacementDirection := displacement.Clone().Normalize()
			massProduct := particle1.Mass * particle2.Mass
			displacementSquared := displacement.LengthSquared()
			accelerationMagnitude := (gravitationalConstant * massProduct) / displacementSquared

			acceleration := displacementDirection.MultiplyScalar(accelerationMagnitude)
			accelerationSum.Add(*acceleration)
		}

		if accelerationSum.Length() > maxAllowedAcceleration {
			accelerationSum.Normalize().MultiplyScalar(maxAllowedAcceleration)
		}

		particle1.Acceleration = accelerationSum
	}
}
