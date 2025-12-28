package ui

import (
	"log"
	Map "swarm-drones-delivery/internal/map"
	"swarm-drones-delivery/internal/simulation"
)

type Game struct {
	Sim *simulation.Simulation
}

func NewGame(mapPath string) *Game {
	sim := simulation.NewSimulation()
	m, err := Map.ReadMap(mapPath)
	if err != nil {
		log.Fatal("Cannot load map")
	}

	sim.Env.Map = m
	return &Game{
		Sim: sim,
	}
}