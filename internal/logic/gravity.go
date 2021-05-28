package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
)

var gravitationalConstant = 1000.0
var maxAllowedAcceleration = 100.0
var particleSizeDiffConstant = 10.0

func MoveAllParticles(gameState *models.GameState) {
	for _, particle := range gameState.Particles {
		particle.Move()
	}
}

func ApplyGravity(gameState *models.GameState) {
	ResetAcceleration(gameState.Particles)

	numParticles := len(gameState.Particles)
	for i := 0; i < numParticles-1; i++ {
		particle1 := gameState.Particles[i]
		for j := i + 1; j < numParticles; j++ {
			particle2 := gameState.Particles[j]
			displacement := particle2.Position.Clone().Subtract(*particle1.Position)
			displacementSquared := displacement.LengthSquared()
			particle1IsMuchBigger := particle1.Mass/particle2.Mass > particleSizeDiffConstant
			particle2IsMuchBigger := particle2.Mass/particle1.Mass > particleSizeDiffConstant

			if !particle1IsMuchBigger {
				displacementDirection1 := displacement.Clone().Normalize()
				accelerationMagnitude1 := (gravitationalConstant * particle2.Mass) / displacementSquared
				acceleration1 := displacementDirection1.MultiplyScalar(accelerationMagnitude1)
				particle1.Acceleration.Add(*acceleration1)
			}

			if !particle2IsMuchBigger {
				displacementDirection2 := displacement.Clone().Normalize().Reverse()
				accelerationMagnitude2 := (gravitationalConstant * particle1.Mass) / displacementSquared
				acceleration2 := displacementDirection2.MultiplyScalar(accelerationMagnitude2)
				particle2.Acceleration.Add(*acceleration2)
			}
		}
	}

	EnforceMaxAcceleration(gameState.Particles)
}

func ResetAcceleration(particles []*models.Particle) {
	for _, particle := range particles {
		particle.Acceleration = &models.Vector{
			X: 0,
			Y: 0,
		}
	}
}

func EnforceMaxAcceleration(particles []*models.Particle) {
	for _, particle := range particles {
		if particle.Acceleration.Length() > maxAllowedAcceleration {
			particle.Acceleration.Normalize().MultiplyScalar(maxAllowedAcceleration)
		}
	}
}
