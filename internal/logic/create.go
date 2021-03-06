package logic

import (
	"math"
	"math/rand"

	"github.com/KeithClinard/go-particle-simulator/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
)

var particleFieldSize int = 10000
var particleFieldCenterMass float64 = 10000.0
var particleFieldDiskStart float64 = 50
var particleFieldDiskDistance float64 = 0.01

func SpawnNewParticle(x, y int, gameState *models.GameState) {
	x1, y1 := ebiten.CursorPosition()
	velX := x1 - x
	velY := y1 - y
	spawnParticle(float64(x), float64(y), float64(velX), float64(velY), gameState)
}

func spawnParticle(x, y, velX, velY float64, gameState *models.GameState) *models.Particle {
	position := &models.Vector{
		X: x,
		Y: y,
	}
	velocity := &models.Vector{
		X: velX,
		Y: velY,
	}
	particle := models.NewParticle(position, velocity)
	gameState.Particles = append(gameState.Particles, particle)
	return particle
}

func spawnPlanet(x, y, velX, velY float64, gameState *models.GameState) *models.Particle {
	position := &models.Vector{
		X: x,
		Y: y,
	}
	velocity := &models.Vector{
		X: velX,
		Y: velY,
	}
	particle := models.NewParticle(position, velocity)
	gameState.Planets = append(gameState.Planets, particle)
	return particle
}

func SpawnParticleField(x, y int, gameState *models.GameState) {
	centerParticle := spawnPlanet(float64(x), float64(y), 0.0, 0.0, gameState)
	centerParticle.Mass = particleFieldCenterMass
	centerParticle.UpdateSize()

	radius := particleFieldDiskStart
	for i := 0; i < particleFieldSize; i++ {
		radius += particleFieldDiskDistance
		angle := rand.Float64() * 2.0 * math.Pi
		position := models.NewVectorFromAngle(angle, radius).Add(*centerParticle.Position)
		orbitalVelocity := math.Sqrt((gravitationalConstant * particleFieldCenterMass) / radius)
		velocity := models.NewVectorFromOrthogonal(angle, orbitalVelocity)
		spawnParticle(position.X, position.Y, velocity.X, velocity.Y, gameState)
	}
}
