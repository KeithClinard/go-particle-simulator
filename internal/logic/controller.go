package logic

import (
	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func HandleUserInputs(gameState *models.GameState) {
	isLeftMousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	// On mouse down
	if isLeftMousePressed && !gameState.Controller.IsLeftMouseDown {
		onLeftMouseDown(gameState)
	}

	if !isLeftMousePressed && gameState.Controller.IsLeftMouseDown {
		onLeftMouseUp(gameState)
	}
}

func onLeftMouseDown(gameState *models.GameState) {
	x, y := ebiten.CursorPosition()
	gameState.Controller.IsLeftMouseDown = true
	gameState.Controller.LeftMouseStartX = x
	gameState.Controller.LeftMouseStartY = y
}

func onLeftMouseUp(gameState *models.GameState) {
	SpawnNewParticle(gameState.Controller.LeftMouseStartX, gameState.Controller.LeftMouseStartY, gameState)
	gameState.Controller.IsLeftMouseDown = false
	gameState.Controller.LeftMouseStartX = 0
	gameState.Controller.LeftMouseStartY = 0
}
