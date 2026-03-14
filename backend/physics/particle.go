package physics

import vec "kepler/vec"

type kinematic struct {
	position vec.Vec
	velocity vec.Vec
}

type Particle struct {
	movement kinematic
	mass float64
}
