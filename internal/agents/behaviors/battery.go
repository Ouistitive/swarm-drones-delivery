package behaviors

import "swarm-drones-delivery/internal/constants"

type Battery struct {
    Total float64
    Level float64
}

func NewBattery() Battery {
	return Battery{
		Total: constants.MAX_BATTERY,
		Level: constants.MAX_BATTERY,
	}
}

func (b *Battery) Consume(amount float64) bool {
    if b.Level < amount {
        b.Level = 0
        return false
    }
    b.Level -= amount
    return true
}

func (b *Battery) IsEmpty() bool {
    return b.Level <= 0
}

func (b *Battery) Ratio() float64 {
    return b.Level / b.Total
}
