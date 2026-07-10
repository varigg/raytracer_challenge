package core

import "math"

const epsilon = 0.00001

func equals(f1, f2 float64) bool {
	return math.Abs(f1-f2) < epsilon
}
