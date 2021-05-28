package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
)

var outOfBoundsBuffer = 100

func DetectOutOfBounds(gameState *models.GameState) {
	width, height := ebiten.WindowSize()
	minXY := 0 - outOfBoundsBuffer
	maxX := width + outOfBoundsBuffer
	maxY := height + outOfBoundsBuffer

	for _, particle := range gameState.Particles {
		intX := int(particle.Position.X)
		intY := int(particle.Position.Y)
		xOutOfBounds := intX < minXY || intX > maxX
		yOutOfBounds := intY < minXY || intY > maxY
		if xOutOfBounds || yOutOfBounds {
			particle.ShouldBeDestroyed = true
		}
	}
}
