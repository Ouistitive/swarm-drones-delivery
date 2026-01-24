package behaviors

import (
	"swarm-drones-delivery/internal/world"
)

type Memory struct {
	addresses map[world.Address]world.Position
}

func NewMemory() Memory {
	return Memory{
		addresses: make(map[world.Address]world.Position),
	}
}

func (m *Memory) AddAddress(add *world.Address, pos *world.Position) {
	if _, exists := m.addresses[*add]; !exists {
		m.addresses[*add] = *pos
	}
}

func (m *Memory) KnowAddress(addr *world.Address) (world.Position, bool) {
	p, e := m.addresses[*addr]
	return p, e
}