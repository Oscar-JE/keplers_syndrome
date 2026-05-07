package physics

import (
	"fmt"
	vec "kepler/vec"
	"testing"
)

func TestFrontalCollision(t *testing.T) {
	v1 := vec.Init([]float64{1, 0})
	v2 := vec.ScalarMultiplication(-1, v1)
	p1 := v2
	p2 := v1
	r := 0.5
	m := 1.0
	particle1 := InitParticle(p1, v1, m, r)
	particle2 := InitParticle(p2, v2, m, r)
	o1, o2 := dualCollision(particle1, particle2, 1)
	expect1 := InitParticle(p1, v2, m, r)
	expect2 := InitParticle(p2, v1, m, r)
	if o1 != expect1.movement || o2 != expect.movement {
		t.Errorf("Error in FrontalCollision")
	}
}
