package models

type GameState struct {
	ShowDebugInfo bool
	Controller    *Controller
	Planets       []*Particle
	Particles     []*Particle
}
