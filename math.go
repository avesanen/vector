package vector

func Dot(a, b Vector) float64 {
	var r float64
	for i := range a {
		r += a[i] * b[i]
	}
	return r
}

func Add(a, b Vector) Vector {
	var v Vector
	for i := 0; i < len(a); i++ {
		v[i] = a[i] + b[i]
	}
	return v
}

func Sub(a, b Vector) Vector {
	var v Vector
	for i := 0; i < len(a); i++ {
		v[i] = a[i] - b[i]
	}
	return v
}
