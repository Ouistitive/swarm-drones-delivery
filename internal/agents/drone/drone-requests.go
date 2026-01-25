package drone

import (
	"math/rand"
	"swarm-drones-delivery/internal/core"
)

func (d *Drone) setDroneStateAndAction(state AgentState, act ActionType) {
	d.state = state
	d.nextAction = act
}

func (d *Drone) getPerceptionData() core.PerceptionData {
	resp := make(chan core.PerceptionData)
	d.requests.PerceptionChan <- core.PerceptionRequest{
		Pos:             d.pos,
		ResponseChannel: resp,
	}
	return <- resp
}

func (d *Drone) getMissions() *core.DeliveryMission {
	missionsChanResponse := make(chan []core.DeliveryMission)
	d.requests.DeliveryMissionsChan <- core.DeliveryMissionsRequest{ResponseChannel: missionsChanResponse}
	m := <-missionsChanResponse
	if len(m) == 0 {
		return nil
	} else {
		return &m[rand.Intn(len(m))]
	}
}

func (d *Drone) getNearestChargingPoint() *core.ChargingMission {
	missionsChanResponse := make(chan core.ChargingMission)
	d.requests.ChargingMissionChan <- core.ChargingMissionRequest{Agt: d, ResponseChannel: missionsChanResponse}
	m := <-missionsChanResponse
	if m.Ok {
		return &m
	}
	return nil
}

func (d *Drone) move() {
	moveChanResponse := make(chan bool)
	d.requests.MoveChan <- core.MoveRequest{Agt: d, ResponseChannel: moveChanResponse}
	<-moveChanResponse
}

func (d *Drone) grab() {
	pickChanResponse := make(chan bool)
	d.requests.PickChan <- core.PickRequest{
		Agt:             d,
		Deliv:           d.deliveryMission.TargetDelivery,
		ResponseChannel: pickChanResponse,
	}
	<-pickChanResponse
}

func (d *Drone) deliver() {
	pickChanResponse := make(chan bool)
	d.requests.DeliverChan <- core.DeliverRequest{
		Agt:             d,
		Deliv:           d.deliveryMission.TargetDelivery,
		ResponseChannel: pickChanResponse,
	}

	res := <-pickChanResponse
	if res {
		d.deliveryMission = nil
	}
}

func (d *Drone) exitCharging() {
	chargingChanResponse := make(chan bool)
	d.requests.ExitChargingChan <- core.ExitChargingRequest{
		Agt:             d,
		ChargingPoint:   d.chargingMission.TargetCharging,
		ResponseChannel: chargingChanResponse,
	}

	res := <-chargingChanResponse
	if res {
		d.chargingMission = nil
	}
}
