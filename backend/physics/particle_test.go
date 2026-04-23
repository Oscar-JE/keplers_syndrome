package physics

import (
	"fmt"
	vec "kepler/vec"
	"testing"
)

func TestFrontalKollision(t *testing.T) {
	// vi kan i vart fakk skriva kollisonen som den ser ut nu
	v1 := vec.Init([]float64{1, 0})
	v2 := vec.ScalarMultiplication(-1, v1)
	p1 := v2
	p2 := v1
	r := 0.5
	m := 1.0
	particle1 := InitParticle(p1, v1, m, r)
	particle2 := InitParticle(p2, v2, m, r)
	o1, o2 := dualCollision(particle1, particle2, 4)
	fmt.Printf("%+v , %+v", o1, o2)
	t.Errorf("there is no test")
}
