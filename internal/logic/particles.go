package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
)

func MoveAllParticles(gameState *models.GameState) {
	gameTick := 1.0 / 60.0
	for _, particle := range gameState.Particles {
		particle.Move(gameTick)
	}
}
