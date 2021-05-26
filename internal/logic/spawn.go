package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func SpawnNewParticle(x, y int, gameState *models.GameState) {
	x1, y1 := ebiten.CursorPosition()
	velX := x1 - x
	velY := y1 - y
	particle := spawnParticle(float64(x), float64(y), float64(velX), float64(velY))
	gameState.Particles = append(gameState.Particles, particle)
}

func spawnParticle(x, y, velX, velY float64) *models.Particle {
	position := &models.Vector{
		X: x,
		Y: y,
	}
	velocity := &models.Vector{
		X: velX,
		Y: velY,
	}
	return models.NewParticle(position, velocity)
}
