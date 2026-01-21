package drone

import (
	"swarm-drones-delivery/internal/core"
)

func (d *Drone) setDroneStateAndAction(state AgentState, act ActionType) {
	d.state = state
	d.nextAction = act
}

func (d *Drone) move() {
	moveChanResponse := make(chan bool)
	d.moveChan <- core.MoveRequest{Agt: d, ResponseChannel: moveChanResponse}
	<- moveChanResponse
}

func (d *Drone) grab() {
	pickChanResponse := make(chan bool)
	d.pickChan <- core.PickRequest{
		Agt: d, 
		Deliv: d.mission.TargetDelivery, 
		ResponseChannel: pickChanResponse,
	}
	<- pickChanResponse
}

func (d *Drone) deliver() {
	pickChanResponse := make(chan bool)
	d.deliverChan <- core.DeliverRequest{
		Agt: d, 
		Deliv: d.mission.TargetDelivery, 
		ResponseChannel: pickChanResponse,
	}
	
	res := <- pickChanResponse
	if res {
		d.mission = nil
	}
}