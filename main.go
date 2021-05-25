package main

import (
	"container/list"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct {
	sprites       *list.List
	showDebugInfo bool
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (game *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (game *Game) Draw(screen *ebiten.Image) {
	if game.showDebugInfo {
		DrawDebugInfo(game, screen)
	}
}

func DrawDebugInfo(game *Game, screen *ebiten.Image) {
	counter := 0
	// Write your game's rendering.
	DrawDebugLineInfo(fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()), screen, &counter)

	// When the "left mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		DrawDebugLineInfo("You're pressing the 'LEFT' mouse button.", screen, &counter)
	}
}

func DrawDebugLineInfo(debugString string, screen *ebiten.Image, counter *int) {
	for i := 0; i < *counter; i++ {
		debugString = "\n" + debugString
	}

	ebitenutil.DebugPrint(screen, debugString)
	*counter = *counter + 1
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func InitializeGameObject(game *Game) {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Particle Simulation")
	ebiten.SetWindowResizable(true)
	game.sprites = list.New()
	game.showDebugInfo = true
}

func main() {
	game := &Game{}
	InitializeGameObject(game)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
