package drone

import (
	"math"
	"math/rand"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/world"
)

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


func (d *Drone) generateTargetPosition() {
	d.targetPos = d.env.Objects()[rand.Intn(len(d.env.Objects()))].Position()
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
