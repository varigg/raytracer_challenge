package core

import (
	"math"
	"strconv"
	"strings"
)

type Tuple struct {
	X, Y, Z, W float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

func NewPoint(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 1)
}

func NewVector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 0)
}

func NewVectorFromString(v string) Tuple {
	x, y, z := stringToCoordinates(v)
	return NewVector(x, y, z)
}

func NewPointFromString(v string) Tuple {
	x, y, z := stringToCoordinates(v)
	return NewPoint(x, y, z)
}

func stringToCoordinates(v string) (float64, float64, float64) {
	coords := strings.Split(v, ",")
	x, _ := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
	y, _ := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
	z, _ := strconv.ParseFloat(strings.TrimSpace(coords[2]), 64)
	return x, y, z
}

func (t Tuple) IsVector() bool { return t.W == 0 }
func (t Tuple) IsPoint() bool  { return t.W == 1 }

func (t Tuple) Add(o Tuple) Tuple {
	return Tuple{t.X + o.X, t.Y + o.Y, t.Z + o.Z, t.W + o.W}
}

func (t Tuple) Subtract(o Tuple) Tuple {
	return Tuple{t.X - o.X, t.Y - o.Y, t.Z - o.Z, t.W - o.W}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t Tuple) Multiply(s float64) Tuple {
	return Tuple{t.X * s, t.Y * s, t.Z * s, t.W * s}
}

func (t Tuple) Divide(s float64) Tuple {
	return Tuple{t.X / s, t.Y / s, t.Z / s, t.W / s}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func (t Tuple) Normalize() Tuple {
	return t.Divide(t.Magnitude())
}

func (t Tuple) Dot(o Tuple) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z + t.W*o.W
}

// Cross only makes sense for vectors (W = 0).
func (t Tuple) Cross(o Tuple) Tuple {
	return NewVector(
		t.Y*o.Z-t.Z*o.Y,
		t.Z*o.X-t.X*o.Z,
		t.X*o.Y-t.Y*o.X,
	)
}

func (t Tuple) Equals(o Tuple) bool {
	return equals(t.X, o.X) && equals(t.Y, o.Y) &&
		equals(t.Z, o.Z) && equals(t.W, o.W)
}

func (t Tuple) Transform(m *Matrix) Tuple {
	return m.MultiplyWithTuple(t)
}

func (t Tuple) Reflect(normal Tuple) Tuple {
	return t.Subtract(normal.Multiply(2).Multiply(t.Dot(normal)))
}
