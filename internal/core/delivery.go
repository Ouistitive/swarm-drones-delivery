package core

import (
	"swarm-drones-delivery/internal/world"
)

type PackageState int

const (
	FREE PackageState = iota
	GRABBED
	DELIVERED
)

type Package struct {
	pos     world.Position
	State   PackageState
	Carrier IAgent
}

func NewDelivery(pos world.Position) *Package {
	return &Package{
		pos:   pos,
		State: FREE,
	}
}

func (d *Package) SetPosition(newPos world.Position) {
	d.pos = newPos
}

func (d *Package) Position() world.Position {
	if d.Carrier != nil && d.State == GRABBED {
		return d.Carrier.Position()
	}
	return d.pos
}

func (d *Package) IsGrabbable() bool {
	return d.State == FREE
}
