package models

import (
	_ "image/png"
	"math"
)

var minParticleSize float64 = 5
var maxParticleSize float64 = 100

type Particle struct {
	Position     *Vector
	Velocity     *Vector
	Acceleration *Vector
	Mass         float64
	Size         float64
}

func NewParticle(position *Vector, velocity *Vector) *Particle {
	particle := &Particle{
		Position: position,
		Velocity: velocity,
		Acceleration: &Vector{
			X: 0,
			Y: 0,
		},
		Mass: 1,
	}
	particle.UpdateSize()
	return particle
}

func (particle *Particle) Move(tick float64) {
	deltaV := particle.Acceleration.Clone().MultiplyScalar(tick)
	particle.Velocity.Add(*deltaV)
	deltaP := particle.Velocity.Clone().MultiplyScalar(tick)
	particle.Position.Add(*deltaP)
}

func (particle *Particle) UpdateSize() {
	particle.Size = math.Sqrt((4*particle.Mass)/math.Pi) + minParticleSize
	if particle.Size > maxParticleSize {
		particle.Size = maxParticleSize
	}
}
