package drone

//go:generate stringer -type=AgentState
type AgentState int

const (
	StateFindingMission AgentState = iota
	StateWandering
	StateMovingToDelivery
	StateMovingToRecharge
	StateMovingToDestination

	StateGrabbing
	StateDelivering
	StateRecharging
)

//go:generate stringer -type=ActionType
type ActionType int

const (
	ActionMove ActionType = iota
	ActionPick
	ActionDeliver
	ActionRecharge
)