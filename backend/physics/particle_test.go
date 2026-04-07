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

func TestClosestPoint(t *testing.T) {
	t.Skip("skriv test för att motbevisa hitta closest point")
}
