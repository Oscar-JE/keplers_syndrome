package kepler

func Dim(v Vec) int {
	return len(v.values)
}

func ScalarMultiplication(scalar float64, v Vec) Vec {
	var rV Vec = Zero(Dim(v))
	for i := range v.values {
		rV.values[i] = scalar * v.values[i]
	}
	return rV
}

func DotProduct(v1 Vec, v2 Vec) float64 {
	demandEqualDimensions(v1, v2, "dot product")
	sum := 0.0
	for i := range v1.values {
		sum += v1.values[i] * v2.values[i]
	}
	return sum
}

func Add(v1 Vec, v2 Vec) Vec {
	demandEqualDimensions(v1, v2, "Vector Addition")
	res := Zero(Dim(v1))
	for i := range res.values {
		res.values[i] = v1.values[i] + v2.values[i]
	}
	return res
}

func Mirror(vector Vec, planeNormal Vec) Vec {
	// hör får vi kontrollra tecken
	// görs med ett enkelt experiment
	nScaling := DotProduct(vector, planeNormal) / DotProduct(planeNormal, planeNormal)
	return Add(vector, ScalarMultiplication(nScaling, planeNormal))
}

func demandEqualDimensions(v1 Vec, v2 Vec, method string) {
	if Dim(v1) != Dim(v2) {
		errorMsg := "Non matching dimensions in method: " + method
		panic(errorMsg)
	}
}
