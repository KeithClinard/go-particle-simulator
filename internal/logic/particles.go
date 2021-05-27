package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
)

var gravityPowerConstant = 1000.0
var maxAllowedAcceleration = 100.0

func MoveAllParticles(gameState *models.GameState) {
	gameTick := 1.0 / 60.0
	for _, particle := range gameState.Particles {
		particle.Move(gameTick)
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
			displacementMagnitude := displacement.Length()
			displacementDirection := displacement.Clone().Normalize()
			accelerationMagnitude := particle2.Mass / (displacementMagnitude * displacementMagnitude)

			acceleration := displacementDirection.MultiplyScalar(accelerationMagnitude)
			accelerationSum.Add(*acceleration)
		}
		accelerationSum.MultiplyScalar(gravityPowerConstant)

		if accelerationSum.Length() > maxAllowedAcceleration {
			accelerationSum.Normalize().MultiplyScalar(maxAllowedAcceleration)
		}

		particle1.Acceleration = accelerationSum
	}
}
