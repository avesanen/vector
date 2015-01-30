package vector

import "testing"

func TestClosestApproach(t *testing.T) {
	Pa := Vector{0, -1}
	Pb := Vector{10, 1}
	Va := Vector{-10, 0}
	Vb := Vector{0, 0}
	Ra := 1.0
	Rb := 2.0
	time, collision := TimeOfClosestApproach(Pa, Pb, Va, Vb, Ra, Rb)
	t.Logf("Closest approach at %v, collision:%v", time, collision)
}
