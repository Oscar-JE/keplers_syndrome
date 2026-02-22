package kepler

import "fmt"

type Vec struct {
	values []float64
}

func Zero(dim int) Vec {
	if dim < 1 {
		panic("In this project we do not allow silly Dimensional vectors")
	}
	var v = make([]float64, dim)
	for i := range v {
		v[i] = 0
	}
	return Vec{values: v}
}

func Init(values []float64) Vec {
	return Vec{values: values}
}

func Eq(a Vec, b Vec) bool {
	if len(a.values) != len(b.values) {
		return false
	}
	var eq bool = true
	for i := range a.values {
		eq = eq && a.values[i] == b.values[i]
	}
	return eq
}

func (v Vec) String() string {
	var rep string = "("
	if len(v.values) > 0 {
		rep += fmt.Sprint(v.values[0])
	}
	for i := 1; i < len(v.values); i++ {
		rep += ", " + fmt.Sprint(v.values[i])
	}
	rep += ")"
	return rep
}
