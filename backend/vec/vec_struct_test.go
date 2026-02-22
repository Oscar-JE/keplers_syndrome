package kepler

import (
	"fmt"
	"testing"
)

func TestPrintDim2(t *testing.T) {
	v := Zero(2)
	rep := fmt.Sprint(v)
	exp := "(0, 0)"
	if rep != exp {
		t.Errorf("expeted %s but was %s", exp, rep)
	}
}

func TestPrintDim1(t *testing.T) {
	v := Zero(1)
	rep := fmt.Sprint(v)
	exp := "(0)"
	if rep != exp {
		t.Errorf("expeted %s but was %s", exp, rep)
	}
}

func TestEqDim(t *testing.T) {
	a := Zero(2)
	b := Zero(1)
	if Eq(a, b) {
		t.Errorf("dimmensioner olika ska inte stämma överens")
	}
}

func TestEqDiffvalues(t *testing.T) {
	a := Zero(3)
	b := Init([]float64{1, 2, 3})
	if Eq(a, b) {
		t.Errorf("olika värden ska ge false")
	}
}

func TestEqTrue(t *testing.T) {
	a := Zero(4)
	b := Zero(4)
	if !Eq(a, b) {
		t.Errorf("noll är lika med noll")
	}
}
