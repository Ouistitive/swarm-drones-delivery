package world

type DeliveryDestination struct {
	Pos Position
	Address Address
}

func NewDeliveryDestination(pos Position, address Address) DeliveryDestination {
	return DeliveryDestination{
		Pos: 		pos,
		Address: 	address,
	}
}