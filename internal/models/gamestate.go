package models

type GameState struct {
	ShowDebugInfo bool
	Controller    *Controller
	Particles     []*Particle
}
