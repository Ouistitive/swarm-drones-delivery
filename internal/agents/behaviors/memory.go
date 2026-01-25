package behaviors

import (
	"fmt"
	"sort"
	"strings"
	"swarm-drones-delivery/internal/world"
)

type Memory struct {
	addresses       map[world.Address]world.Position
	addressesString string
}

func NewMemory() Memory {
	return Memory{
		addresses: make(map[world.Address]world.Position),
	}
}

func (m *Memory) AddAddress(add *world.Address, pos *world.Position) {
	if _, exists := m.addresses[*add]; !exists {
		m.addresses[*add] = *pos
		m.addAddressString()
	}
}

func (m *Memory) KnowAddress(addr *world.Address) (world.Position, bool) {
	p, e := m.addresses[*addr]
	return p, e
}

func (m *Memory) ToString() string {
	return m.addressesString
}

func (m *Memory) addAddressString() {
	type entry struct {
		street string
		pos    world.Position
	}

	entries := make([]entry, 0, len(m.addresses))
	for k, v := range m.addresses {
		entries = append(entries, entry{
			street: k.Street,
			pos:    v,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].street < entries[j].street
	})

	var b strings.Builder
	for _, e := range entries {
		fmt.Fprintf(&b, "\n - %s (%.0f,%.0f)", e.street, e.pos.X, e.pos.Y)
	}

	m.addressesString = b.String()
}
