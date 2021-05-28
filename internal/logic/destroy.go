package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
)

func DestroyParticles(gameState *models.GameState) {
	oldParticleList := gameState.Particles
	newParticleList := oldParticleList[:0]

	for _, particle := range oldParticleList {
		if !(particle.ShouldBeDestroyed) {
			newParticleList = append(newParticleList, particle)
		}
	}
	for i := len(newParticleList); i < len(oldParticleList); i++ {
		oldParticleList[i] = nil
	}
	gameState.Particles = newParticleList
}
