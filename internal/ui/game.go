package ui

import (
	"log"
	"swarm-drones-delivery/internal/simulation"
	"swarm-drones-delivery/internal/world"
)

type Game struct {
	Sim *simulation.Simulation
}

func NewGame(mapPath string) *Game {
	m, err := world.ReadMap(mapPath)
	sim := simulation.NewSimulation(m)
	if err != nil {
		log.Fatal("Cannot load map")
	}

	g := &Game{
		Sim: sim,
	}

	return g
}
