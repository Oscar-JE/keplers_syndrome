package kepler

import "testing"

func TestScalarMultiplicationZeroVec(t *testing.T) {
	a := Zero(3)
	res := ScalarMultiplication(100, a)
	if !Eq(res, Zero(3)) {
		t.Errorf("scaling a zero vector should result in a zero vector")
	}
}

func TestScalarMultiplicationZeroScalar(t *testing.T) {
	a := Init([]float64{5})
	res := ScalarMultiplication(0, a)
	if !Eq(res, Zero(1)) {
		t.Errorf("a scaling of zero should result in zero")
	}
}

func TestDoubleTheLenght(t *testing.T) {
	a := Init([]float64{5})
	res := ScalarMultiplication(2, a)
	expected := Init([]float64{10})
	if !Eq(res, expected) {
		t.Errorf("a scaling of two should double the vector")
	}
}

func TestDotProduct90(t *testing.T) {
	a := Init([]float64{1, 1})
	b := Init([]float64{1, -1})
	if DotProduct(a, b) != 0 {
		t.Errorf("orthogonal vectors have a zero dotproduct between each other")
	}
}

func TestDotProduct0(t *testing.T) {
	a := Init([]float64{2, 0})
	if DotProduct(a, a) != 4 {
		t.Errorf("Dot product of parallel is the lengths multiplied")
	}
}
