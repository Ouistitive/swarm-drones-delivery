package simulation

type Simulation struct {
	Env *Environment
}

func NewSimulation() (*Simulation) {
	return &Simulation{}
}