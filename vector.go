package vector

import (
	"math"
)

// Vector is a [2]float64 type with extra methods for calculating and
// manipulating length and angle.
type Vector [2]float64

// Normalize sets the length of vector to 1.
func (v *Vector) Normalize() {
	l := v.Length()
	for i := 0; i < len(v); i++ {
		v[i] = v[i] / l
	}
}

// Length returns the vector's length.
func (v *Vector) Length() float64 {
	var sum float64
	for i := 0; i < len(v); i++ {
		sum += math.Pow(v[i], 2)
	}
	return math.Sqrt(sum)
}

// SetLength sets the length of vector so that the angle stays the same.
func (v *Vector) SetLength(l float64) {
	v.Normalize()
	for i := 0; i < len(v); i++ {
		v[i] *= l
	}
}

// SetAngle sets the angle of vector so that the length stays the same.
func (v *Vector) SetAngle(angle float64) {
	l := v.Length()
	v[0] = l * math.Cos(angle)
	v[1] = l * math.Sin(angle)
}

// Rotate rotates the vector, keeping length the same.
func (v *Vector) Rotate(angle float64) {
	v.SetAngle(v.Angle() + angle)
}

// Angle returns angle in radians.
func (v *Vector) Angle() float64 {
	return math.Atan2(v[1], v[0])
}

func CircleCollision(vPos, vVel, cPos Vector, cRad float64) []Vector {
	E := vPos
	L := vVel
	C := cPos
	r := cRad
	d := Sub(L, E)
	f := Sub(E, C)
	a := Dot(d, d)
	b := 2 * Dot(f, d)
	c := Dot(f, f) - r*r

	disc := b*b - 4*a*c

	if disc < 0 {
		// No intersection
		return nil
	}

	disc = math.Sqrt(disc)

	t1 := (-b - disc) / (2 * a)
	t2 := (-b + disc) / (2 * a)

	c1 := Vector{vPos[0] + d[0]*t1, vPos[1] + d[1]*t1}
	c2 := Vector{vPos[0] + d[0]*t2, vPos[1] + d[1]*t2}

	t1hit := t1 >= 0 && t1 <= 1
	t2hit := t2 >= 0 && t2 <= 1

	if t1hit && t2hit {
		return []Vector{c1, c2}
	}

	if t1hit && !t2hit {
		return []Vector{c1}
	}

	if !t1hit && t2hit {
		return []Vector{c2}
	}

	return nil
}
