package core

type MoveRequest struct {
	Agt IAgent
	ResponseChannel chan bool	
}

type SpawnRequest struct {
	Agt IAgent
	ResponseChannel chan bool
}