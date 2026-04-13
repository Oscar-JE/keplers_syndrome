package physics

import (
	"fmt"
	vec "kepler/vec"
	"testing"
)

func TestFrontalKollision(t *testing.T) {
	t.Skip("sparar kilisiont till lite längre fram")
	v1 := vec.Init([]float64{1, 0})
	v2 := vec.ScalarMultiplication(-1, v1)
	fmt.Println(v2)
}

func TestClosestPointEnterOrigo(t *testing.T) {
	start := vec.Init([]float64{-5, 0})
	velocity := vec.Init([]float64{1, 0})
	closest := closestPointToZero(start, velocity)
	if !vec.Eq(closest, vec.Zero(2)) {
		t.Errorf("this line should clearly pass through origo")
	}
}

func TestClosestPointAboveOrigo(t *testing.T) {
	start := vec.Init([]float64{-5, 5})
	velocity := vec.Init([]float64{1, 0})
	closest := closestPointToZero(start, velocity)
	expected := vec.Init([]float64{0, 5})
	if !vec.Eq(closest, expected) {
		t.Errorf("this line should clearly pass through origo")
	}
}

func TestClosestPointTwisted(t *testing.T) {
	start := vec.Init([]float64{-2, 0})
	velocity := vec.Init([]float64{1, 1})
	closest := closestPointToZero(start, velocity)
	expected := vec.Init([]float64{-1, 1})
	if !vec.Eq(closest, expected) {
		t.Errorf("Inspect case closestPointTwisted")
	}
}
