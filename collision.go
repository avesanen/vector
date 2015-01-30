package vector

import (
	"math"
)

// Given two circles with position (Pa,Pb), velocity (Va,Vb) and radius (Ra, Rb),
// Return the time of closest approach and possible collision (float64, bool).
func TimeOfClosestApproach(Pa, Pb, Va, Vb Vector, Ra float64, Rb float64) (float64, bool) {
	var collision bool
	var Pab Vector = Sub(Pa, Pb)
	var Vab Vector = Sub(Va, Vb)
	a := Dot(Vab, Vab)
	b := 2 * Dot(Vab, Vab)
	c := Dot(Pab, Pab) - (Ra+Rb)*(Ra+Rb)
	discriminant := b*b - 4*a*c
	var t float64
	if discriminant < 0 {
		t = -b / 2 * a
		collision = false
	} else {
		t0 := (-b + math.Sqrt(discriminant)/(2*a))
		t1 := (-b - math.Sqrt(discriminant)/(2*a))
		t = math.Min(t0, t1)

		if t < 0 {
			collision = false
		} else {
			collision = true
		}
	}
	if t < 0 {
		t = 0
	}
	return t, collision
}
