package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
)

var gravitationalConstant = 1000.0
var maxAllowedAcceleration = 100.0
var outOfBoundsBuffer = 100
var particleSizeDiffConstant = 10.0

func MoveAllParticles(gameState *models.GameState) {
	for _, particle := range gameState.Particles {
		particle.Move()
	}
}

func ApplyGravity(gameState *models.GameState) {
	for _, particle := range gameState.Particles {
		particle.Acceleration = &models.Vector{
			X: 0,
			Y: 0,
		}
	}

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

	for _, particle := range gameState.Particles {
		if particle.Acceleration.Length() > maxAllowedAcceleration {
			particle.Acceleration.Normalize().MultiplyScalar(maxAllowedAcceleration)
		}
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
