package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
)

var gravitationalConstant = 1000.0
var maxAllowedAcceleration = 100.0

func ApplyGravity(gameState *models.GameState) {
	ResetAcceleration(gameState.Particles)

	for _, particle := range gameState.Particles {
		for _, planet := range gameState.Planets {
			displacement := planet.Position.Clone().Subtract(*particle.Position)
			displacementSquared := displacement.LengthSquared()
			displacementDirection := displacement.Clone().Normalize()
			accelerationMagnitude := (gravitationalConstant * planet.Mass) / displacementSquared
			acceleration1 := displacementDirection.MultiplyScalar(accelerationMagnitude)
			particle.Acceleration.Add(*acceleration1)
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
