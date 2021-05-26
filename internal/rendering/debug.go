package rendering

import (
	"fmt"

	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawDebugInfo(gameState *models.GameState, screen *ebiten.Image) {
	if !gameState.ShowDebugInfo {
		return
	}
	counter := 0
	// Write your game's rendering.
	drawDebugLineInfo(fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()), screen, &counter)

	// When the "left mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		drawDebugLineInfo("You're pressing the 'LEFT' mouse button.", screen, &counter)
	}
}

func drawDebugLineInfo(debugString string, screen *ebiten.Image, counter *int) {
	for i := 0; i < *counter; i++ {
		debugString = "\n" + debugString
	}

	ebitenutil.DebugPrint(screen, debugString)
	*counter = *counter + 1
}
