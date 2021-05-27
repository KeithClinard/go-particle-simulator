package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
)

var gravitationalConstant = 1000.0
var maxAllowedAcceleration = 100.0
var outOfBoundsBuffer = 100

func MoveAllParticles(gameState *models.GameState) {
	for _, particle := range gameState.Particles {
		particle.Move()
	}
}

func ApplyGravity(gameState *models.GameState) {
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

func DestroyOutOfBounds(gameState *models.GameState) {
	width, height := ebiten.WindowSize()
	minXY := 0 - outOfBoundsBuffer
	maxX := width + outOfBoundsBuffer
	maxY := height + outOfBoundsBuffer

	oldParticleList := gameState.Particles
	newParticleList := oldParticleList[:0]

	for _, particle := range oldParticleList {
		intX := int(particle.Position.X)
		intY := int(particle.Position.Y)
		xOutOfBounds := intX < minXY || intX > maxX
		yOutOfBounds := intY < minXY || intY > maxY
		if !(xOutOfBounds || yOutOfBounds) {
			newParticleList = append(newParticleList, particle)
		}
	}
	for i := len(newParticleList); i < len(oldParticleList); i++ {
		oldParticleList[i] = nil
	}
	gameState.Particles = newParticleList
}
