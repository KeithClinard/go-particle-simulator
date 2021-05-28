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
	isRightMousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)

	// On mouse down
	if isRightMousePressed && !gameState.Controller.IsRightMouseDown {
		onRightMouseDown(gameState)
	}

	if !isRightMousePressed && gameState.Controller.IsRightMouseDown {
		onRightMouseUp(gameState)
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

func onRightMouseDown(gameState *models.GameState) {
	x, y := ebiten.CursorPosition()
	gameState.Controller.IsRightMouseDown = true
	gameState.Controller.RightMouseStartX = x
	gameState.Controller.RightMouseStartY = y
}

func onRightMouseUp(gameState *models.GameState) {
	SpawnParticleField(gameState.Controller.RightMouseStartX, gameState.Controller.RightMouseStartY, gameState)
	gameState.Controller.IsRightMouseDown = false
	gameState.Controller.RightMouseStartX = 0
	gameState.Controller.RightMouseStartY = 0
}
