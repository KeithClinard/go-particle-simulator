package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
)

func DestroyParticles(gameState *models.GameState) {
	gameState.Particles = destroyParticlesInList(gameState.Particles)
	gameState.Planets = destroyParticlesInList(gameState.Planets)
}

func destroyParticlesInList(oldParticleList []*models.Particle) []*models.Particle {
	newParticleList := oldParticleList[:0]

	for _, particle := range oldParticleList {
		if !(particle.ShouldBeDestroyed) {
			newParticleList = append(newParticleList, particle)
		}
	}
	for i := len(newParticleList); i < len(oldParticleList); i++ {
		oldParticleList[i] = nil
	}
	return newParticleList
}
