package world

type ChargingPointState int

const (
	ChargingPointFree ChargingPointState = iota
	ChargingPointOccupied
)

type ChargingPoint struct {
	Pos 	Position
	State 	ChargingPointState
}

func NewChargingPoint(pos Position) ChargingPoint {
	return ChargingPoint{
		Pos: 	pos,
		State: 	ChargingPointFree,
	}
}