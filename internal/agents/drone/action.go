package drone

import "swarm-drones-delivery/internal/constants"

func (d *Drone) Act() {
	switch d.nextAction {
	case ActionMove:
		if d.battery.Consume(constants.BATTERY_DISCHARGING_MOVE) {
			d.move()
		}
	case ActionPick:
		if d.battery.Consume(constants.BATTERY_DISCHARGING_PICK) {
			d.grab()
		}
	case ActionDeliver:
		if d.battery.Consume(constants.BATTERY_DISCHARGING_DELIVER) {
			d.deliver()
		}
	case ActionRecharge:
		d.battery.Recharge(constants.BATTERY_CHARGING_RATE)
		if d.battery.IsFull() {
			d.deliveryMission = nil
			d.exitCharging()
			d.setDroneStateAndAction(StateFindingMission, ActionMove)
		}
	}
}
