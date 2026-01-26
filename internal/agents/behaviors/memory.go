package behaviors

import (
	"fmt"
	"strings"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"
)

type Memory struct {
	addressesPosition []world.DeliveryDestination
	addressesString string
}

func NewMemory() Memory {
	return Memory{
		addressesPosition: make([]world.DeliveryDestination, 0),
	}
}

func (m *Memory) AddAddress(add *world.Address, pos *world.Position) {
	if _, exists := m.KnowAddress(add); !exists {
		m.addressesPosition = append(m.addressesPosition, world.NewDeliveryDestination(*pos, *add))
		m.addAddressString()
	}
}

func (m *Memory) KnowAddress(newAddr *world.Address) (world.Position, bool) {
	for _, addr := range m.addressesPosition {
		if addr.Address == *newAddr {
			return addr.Pos, true
		}
	}
	return world.NullPosition(), false
}

func (m *Memory) ToString() string {
	return m.addressesString
}

func (m *Memory) addAddressString() {
	var b strings.Builder
	for _, e := range m.addressesPosition {
		fmt.Fprintf(&b, "\n - %s (%s)", e.Address.Street, utils.PositionToString(e.Pos))
	}

	m.addressesString = b.String()
}
