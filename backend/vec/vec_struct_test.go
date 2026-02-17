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
