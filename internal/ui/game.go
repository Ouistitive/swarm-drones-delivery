package ui

import "swarm-drones-delivery/internal/simulation"

type Game struct {
	simulation *simulation.Simulation
}

func NewGame() *Game {
	return &Game{
		simulation: simulation.NewSimulation(),
	}
}
