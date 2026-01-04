package core

type MoveRequest struct {
	Agt IAgent
	ResponseChannel chan bool	
}