package models

import (
	_ "image/png"
	"math"
)

var minParticleSize float64 = 5
var maxParticleSize float64 = 100
var tick float64 = 1.0 / 60.0

type Particle struct {
	Position          *Vector
	Velocity          *Vector
	Acceleration      *Vector
	Mass              float64
	Size              float64
	ShouldBeDestroyed bool
}

func NewParticle(position *Vector, velocity *Vector) *Particle {
	particle := &Particle{
		Position: position,
		Velocity: velocity,
		Acceleration: &Vector{
			X: 0,
			Y: 0,
		},
		Mass:              1,
		ShouldBeDestroyed: false,
	}
	particle.UpdateSize()
	return particle
}

func (particle *Particle) Move() {
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

func (particle *Particle) Absorb(other *Particle) {
	sumMass := particle.Mass + other.Mass
	weightedPosition1 := particle.Position.MultiplyScalar(particle.Mass)
	weightedPosition2 := other.Position.MultiplyScalar(other.Mass)
	particle.Position = weightedPosition1.Clone().Add(*weightedPosition2).DivideScalar(sumMass)

	weightedVelocity1 := particle.Velocity.MultiplyScalar(particle.Mass)
	weightedVelocity2 := other.Velocity.MultiplyScalar(other.Mass)
	particle.Velocity = weightedVelocity1.Clone().Add(*weightedVelocity2).DivideScalar(sumMass)

	particle.Mass = sumMass
	other.ShouldBeDestroyed = true
}
