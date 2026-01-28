package drone

import "fmt"

func (d *Drone) GetDisplayData() string {
	mission := "None"
	if d.deliveryMission != nil {
		mission = d.deliveryMission.ToString()
	} else if d.chargingMission != nil {
		mission = d.chargingMission.ToString()
	}

	text := fmt.Sprintf(
		"AgentID: %s\nState: %s\nAction: %s\nBattery: %d%%\nMemory: %s\nMission: %s",
		d.id,
		d.state.String(),
		d.nextAction.String(),
		int(d.battery.Ratio()*100),
		d.memory.ToString(),
		mission,
	)

	return text
}