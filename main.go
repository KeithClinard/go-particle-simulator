package main

import (
	"log"

	"github.com/KeithClinard/go-particle-simulator/internal/logic"
	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/KeithClinard/go-particle-simulator/internal/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	gameState *models.GameState
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (game *Game) Update() error {
	logic.HandleUserInputs(game.gameState)
	logic.ApplyParticleGravity(game.gameState)
	logic.MoveAllParticles(game.gameState)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (game *Game) Draw(screen *ebiten.Image) {
	rendering.DrawParticles(game.gameState, screen)
	rendering.DrawDebugInfo(game.gameState, screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
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
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
