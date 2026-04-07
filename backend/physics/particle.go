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

func timeToCollision(pos1 vec.Vec, v1 vec.Vec, pos2 vec.Vec, v2 vec.Vec, distance float64) float64 {
	deltaPos := vec.Subtract(pos2, pos1)
	deltaV := vec.Subtract(v2, v1)
	if vec.DotProduct(deltaPos, deltaPos) < distance*distance {
		return 0
	}
	closestPoint := closestPoint(deltaPos, deltaV)

	return vec.DotProduct(closestPoint, closestPoint) // vet att detta är fel
}

func closestPoint(deltaPos vec.Vec, deltaV vec.Vec) vec.Vec {
	fromClosestPoint := vec.ScalarMultiplication(vec.DotProduct(deltaV, deltaPos)/vec.DotProduct(deltaV, deltaV), deltaV)
	return vec.Subtract(deltaPos, fromClosestPoint)
}

func dualCollision(p1 Particle, p2 Particle, timeStep float64) (Particle, Particle) {
	return Particle{}, Particle{}
}
