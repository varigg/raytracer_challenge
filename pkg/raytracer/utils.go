package raytracer

import "math"

const EPSILON = 0.00001

func equals(f1, f2 float64) bool {
	if math.Abs(f1-f2) < EPSILON {
		return true
	}
	return false
}
