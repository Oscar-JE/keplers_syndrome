package physics

import vec "kepler/vec"

type kinematic struct {
	position vec.Vec
	velocity vec.Vec
}

type Particle struct {
	movement kinematic
	mass     float64
	radia    float64
}

type kollisionDetailjs struct {
	velocity vec.Vec
	mass     float64
}

func kollisionDetailjsFromParticle(p Particle) kollisionDetailjs {
	return kollisionDetailjs{velocity: p.movement.velocity, mass: p.mass}
}

func timeToCollision(pos1 vec.Vec, v1 vec.Vec, pos2 vec.Vec, v2 vec.Vec, distance float64) (bool, float64) {
	deltaPos := vec.Subtract(pos2, pos1)
	deltaV := vec.Subtract(v2, v1)
	if vec.DotProduct(deltaPos, deltaPos) < distance*distance {
		return true, 0
	}
	closestPoint := closestPointToZero(deltaPos, deltaV)
	// börja med att kolla om de koliderar
	//borde vara att kolla om det närmaste avståndet är

	if vec.DotProduct(closestPoint, closestPoint) > distance*distance { // det blir ingen kollision
		return false, -1
	} else {

	}

	return true, vec.DotProduct(closestPoint, closestPoint)
}

func closestPointToZero(startPosition vec.Vec, velocity vec.Vec) vec.Vec {
	fromClosestPoint := vec.ScalarMultiplication(vec.DotProduct(velocity, startPosition)/vec.DotProduct(velocity, velocity), velocity)
	return vec.Subtract(startPosition, fromClosestPoint)
}

func dualCollision(p1 Particle, p2 Particle, timeStep float64) (Particle, Particle) {
	return Particle{}, Particle{}
}
