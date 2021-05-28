package main

import (
	"log"

	"github.com/KeithClinard/go-particle-simulator/internal/logic"
	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/KeithClinard/go-particle-simulator/internal/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

var tickCounter uint64 = 0

// Only calculate gravity every N ticks
// Particles still move between calculations
// Increase to improve performance at the cost of accuracy
var gravityFrequencyConstant uint64 = 1

type Game struct {
	gameState *models.GameState
}

func (game *Game) Update() error {
	logic.HandleUserInputs(game.gameState)
	if tickCounter%gravityFrequencyConstant == 0 {
		logic.ApplyGravity(game.gameState)
	}
	logic.MoveAllParticles(game.gameState)
	logic.DestroyOutOfBounds(game.gameState)
	tickCounter++
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	rendering.DrawParticles(game.gameState, screen)
	rendering.DrawDebugInfo(game.gameState, screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func InitializeGameObject() *Game {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Particle Simulation")
	ebiten.SetWindowResizable(true)
	return &Game{
		gameState: &models.GameState{
			ShowDebugInfo: true,
			Particles:     make([]*models.Particle, 0),
			Controller:    new(models.Controller),
		},
	}
}

func main() {
	game := InitializeGameObject()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
