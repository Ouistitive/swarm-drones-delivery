package agents

import (
	"fmt"
	"math"
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

type Drone struct {
	id  core.AgentID
	env core.IEnvironment

	vision          behaviors.Vision
	surroundingAgts []core.IAgent

	syncChan         chan int
	moveChan         chan core.MoveRequest
	moveChanResponse chan bool

	pos 			world.Position
	targetPos		world.Position
	targetDir 		world.Position
	currentDir 		world.Position
	velocity 		float64

	t time.Time

}

func (d *Drone) SurroundingAgents() []core.IAgent {
	return d.surroundingAgts
}

func DroneFactory(
	env core.IEnvironment,
	agtId core.AgentID,
	syncChan chan int,
) core.AgentFactory {
	return func(pos world.Position, moveChan chan core.MoveRequest) core.IAgent {
		return NewDrone(
			env,
			agtId,
			pos,
			syncChan,
			moveChan,
		)
	}
}

func (d *Drone) ID() core.AgentID {
	return d.id
}

func (d *Drone) Start() {
	fmt.Println("Drone started:", d.id)

	for {
		step := <-d.syncChan
		d.Percept()
		d.Deliberate()
		d.Act()
		d.syncChan <- step + 1
	}
}

func (d *Drone) Percept() {
	agts := d.env.Agents()
	d.surroundingAgts = d.surroundingAgts[:0]

	for _, a := range agts {
		if d.vision.IsAgentDetected(d.pos, a) {
			d.surroundingAgts = append(d.surroundingAgts, a)
		}
	}
}

func (d *Drone) Deliberate() {
	if time.Since(d.t) >= 3 * time.Second {
		d.generateTargetPosition()
		d.changeTargetAngle()
		
		d.t = time.Now()
	}
}

func (d *Drone) Act() {
	moveChanResponse := make(chan bool)
	d.moveChan <- core.MoveRequest{Agt: d, ResponseChannel: moveChanResponse}
	<- moveChanResponse
}

func (d *Drone) Move() {
	dir, distance := d.vectorToTarget()

	// If the drone is close enough to the target, the drone is "glued" to the target pos
	if distance <= constants.CLOSE_TO_TARGET {
		d.pos = d.targetPos
		d.velocity = 0
		return
	}

	d.targetDir = dir
	if d.currentDir.X < d.targetDir.X {
		d.currentDir.X += constants.INCREMENTAL_DIRECTION
	} else if d.currentDir.X > d.targetDir.X {
		d.currentDir.X -= constants.INCREMENTAL_DIRECTION
	}

	if d.currentDir.Y < d.targetDir.Y {
		d.currentDir.Y += constants.INCREMENTAL_DIRECTION
	} else if d.currentDir.Y > d.targetDir.Y {
		d.currentDir.Y -= constants.INCREMENTAL_DIRECTION
	}

	d.adjustVelocity(distance)
	d.pos.X += d.currentDir.X * d.velocity
	d.pos.Y += d.currentDir.Y * d.velocity
}

func NewDrone(env core.IEnvironment, agtId core.AgentID, pos world.Position, syncChan chan int, moveChan chan core.MoveRequest) *Drone {
	return &Drone{
		t: time.Now(),
		env:              env,
		id:               agtId,
		vision:           behaviors.NewVision(constants.VISION_RANGE),
		syncChan:         syncChan,
		moveChan:         moveChan,
		pos:              pos,
		moveChanResponse: make(chan bool),
		surroundingAgts:  []core.IAgent{},
		targetDir: 		  world.NullPosition(),
		currentDir: 	  world.NullPosition(),
		targetPos: 		  env.World().RandomPosition(),
		velocity: 		  0.0,
	}
}

func (d *Drone) Position() world.Position {
	return d.pos
}

func (d *Drone) generateTargetPosition() {
	d.targetPos = d.env.World().RandomPosition()
}

func (d *Drone) adjustVelocity(distance float64) {
	targetVelocity := constants.MAX_VELOCITY

	if distance < constants.LOW_DISTANCE {
		targetVelocity = math.Max(
			constants.MIN_VELOCITY,
			constants.MAX_VELOCITY * (distance / constants.LOW_DISTANCE),
		)
	}

	if d.velocity < targetVelocity {
		d.velocity = math.Min(
			targetVelocity,
			d.velocity + constants.INCREMENTAL_VELOCITY,
		)
	} else {
		d.velocity = targetVelocity
	}
}

func (d *Drone) changeTargetAngle() {
	d.targetDir, _ = d.vectorToTarget()
	// If the drone is slow enough, the direction goes automatically to the target dir 
	if d.velocity < constants.SLOW_VELOCITY_THRESHOLD {
		d.currentDir = d.targetDir
	}
}

func (d *Drone) TargetPos() world.Position {
	return d.targetPos
}

func (d *Drone) vectorToTarget() (dir world.Position, distance float64) {
	dx := d.targetPos.X - d.pos.X
	dy := d.targetPos.Y - d.pos.Y

	distance = math.Hypot(dx, dy)
	if distance <= constants.CLOSE_TO_TARGET {
		return world.NullPosition(), 0
	}
	
	dir = world.Position{X: dx / distance, Y: dy / distance}
	return
}
