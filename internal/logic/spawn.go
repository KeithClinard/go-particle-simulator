package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func SpawnNewParticle(gameState *models.GameState) {
	wasLeftClickDown := gameState.Controller.IsLeftMouseDown
	isLeftClickReleased := !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if wasLeftClickDown && isLeftClickReleased {
		spawnParticleFromState(gameState)
		gameState.Controller.IsLeftMouseDown = false
		gameState.Controller.LeftMouseStartX = 0
		gameState.Controller.LeftMouseStartY = 0
	}
}

func spawnParticleFromState(gameState *models.GameState) {
	x := gameState.Controller.LeftMouseStartX
	y := gameState.Controller.LeftMouseStartY
	x1, y1 := ebiten.CursorPosition()
	velX := x1 - x
	velY := y1 - y
	spawnParticle(x, y, velX, velY)
}

func spawnParticle(x, y, velX, velY int) {
	// TODO
}
