package drone

func (d *Drone) Percept() {
	data := d.getPerceptionData()
	d.allDeliveryMissions = data.Missions
	d.surroundingAgts = d.surroundingAgts[:0]

	for _, a := range data.Agents {
		if d.vision.IsAgentDetected(&d.pos, &a.Pos) {
			d.surroundingAgts = append(d.surroundingAgts, a)
		}
	}

	for _, de := range data.Destinations {
		if d.vision.IsAgentDetected(&d.pos, &de.Pos) {
			d.memory.AddAddress(&de.Address, &de.Pos)
		}
	}
}

func (d *Drone) ProcessInbox() {
	for {
		select {
		case msg := <- d.inbox:
			for _, addr := range msg.Addresses {
				d.memory.AddAddress(&addr.Address, &addr.Pos)
			}
		default:
			return
		}
	}
}