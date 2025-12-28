package simulation

type Simulation struct {
	env *Environment
}

func NewSimulation() (*Simulation) {
	return &Simulation{}
}