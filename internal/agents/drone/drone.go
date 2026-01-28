package drone

import (
	"fmt"
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

type Drone struct {
	id         core.AgentID
	hasSpawned bool

	vision          behaviors.Vision
	surroundingAgts []core.AgentView

	syncChan chan int
	requests core.ChannelRequests
	inbox    chan DroneSharedData

	pos        world.Position
	targetPos  world.Position // What the drone is trying to go in a current state
	targetDir  world.Position
	currentDir world.Position
	velocity   float64

	state      AgentState
	nextAction ActionType

	scoredMissions 		[]core.ScoredMission
	allDeliveryMissions	[]core.DeliveryMission
	deliveryMission 	*core.DeliveryMission
	chargingMission 	*core.ChargingMission
	battery         	behaviors.Battery
	memory          	behaviors.Memory

	t time.Time
}

func (d *Drone) Spawned() bool {
	return d.hasSpawned
}

func (d *Drone) SurroundingAgents() []core.AgentView {
	return d.surroundingAgts
}

func (d *Drone) ID() core.AgentID {
	return d.id
}

func (d *Drone) Position() world.Position {
	return d.pos
}

func (d *Drone) TargetPos() world.Position {
	return d.targetPos
}

func (d *Drone) Mission() *core.DeliveryMission {
	return d.deliveryMission
}

func (d *Drone) Start() {
	fmt.Println("Drone started:", d.id)

	for !d.hasSpawned {
		startChanResponse := make(chan bool)
		d.requests.SpawnChan <- core.SpawnRequest{Agt: d, ResponseChannel: startChanResponse}
		d.hasSpawned = <-startChanResponse
	}

	for {
		step := <-d.syncChan
		d.Percept()
		d.ProcessInbox()
		d.Deliberate()
		d.Act()
		d.syncChan <- step + 1
	}
}
