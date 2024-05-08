package raytracer

import (
	"math"
	"strconv"
	"strings"
)

// type Tuple interface {
// 	X() float64
// 	Y() float64
// 	Z() float64
// 	W() float64
// 	IsVector() bool
// 	IsPoint() bool
// 	Add(Tuple) Tuple
// 	Subtract(Tuple) Tuple
// 	Negate() Tuple
// 	Multiply(float64) Tuple
// 	Divide(float64) Tuple
// }

// 	Divide(float64) Tuple
// 	Magnitude() float64
// 	Normalize() Tuple
// 	Dot(Tuple) float64
// 	Cross(Tuple) Tuple
// }

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func NewTuple(a, b, c, w float64) *Tuple {
	t := &Tuple{
		x: a,
		y: b,
		z: c,
		w: w,
	}
	return t

}

func NewPoint(a, b, c float64) *Tuple {
	return NewTuple(a, b, c, 1.0)
}

func NewVector(a, b, c float64) *Tuple {
	return NewTuple(a, b, c, 0.0)
}

func (t *Tuple) X() float64 {
	return t.x
}

func (t *Tuple) Y() float64 {
	return t.y
}

func (t *Tuple) Z() float64 {
	return t.z
}

func (t *Tuple) W() float64 {
	return t.w
}

func NewVectorFromString(v string) *Tuple {
	x, y, z := stringToCoordinates(v)
	return NewVector(x, y, z)
}

func NewPointFromString(v string) *Tuple {
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

func (t *Tuple) IsVector() bool {
	return t.w == 0.0
}

func (t *Tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t1 *Tuple) Add(t2 *Tuple) *Tuple {
	return &Tuple{
		x: t1.x + t2.X(),
		y: t1.y + t2.Y(),
		z: t1.z + t2.Z(),
		w: t1.w + t2.W(),
	}
}

func (t1 *Tuple) Subtract(t2 *Tuple) *Tuple {
	return &Tuple{
		x: t1.x - t2.X(),
		y: t1.y - t2.Y(),
		z: t1.z - t2.Z(),
		w: t1.w - t2.W(),
	}
}

func (t *Tuple) Negate() *Tuple {
	return &Tuple{
		x: 0 - t.x,
		y: 0 - t.y,
		z: 0 - t.z,
		w: 0 - t.w,
	}
}

func (t *Tuple) Multiply(s float64) *Tuple {
	return &Tuple{
		x: t.x * s,
		y: t.y * s,
		z: t.z * s,
		w: t.w * s,
	}
}

func (t *Tuple) Divide(s float64) *Tuple {
	return &Tuple{
		x: t.x / s,
		y: t.y / s,
		z: t.z / s,
		w: t.w / s,
	}
}

func (v *Tuple) Magnitude() float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z() + v.W()*v.W())
}

func (v *Tuple) Normalize() *Tuple {
	t := NewTuple(v.X()/v.Magnitude(),
		v.Y()/v.Magnitude(),
		v.Z()/v.Magnitude(),
		v.W()/v.Magnitude(),
	)
	return t
}

func (t1 *Tuple) Dot(t2 *Tuple) float64 {
	return t1.X()*t2.X() + t1.Y()*t2.Y() +
		t1.Z()*t2.Z() + t1.W()*t2.W()
}

func (v1 *Tuple) Cross(v2 *Tuple) *Tuple {
	// This makes only sense for vectors
	return NewVector(v1.Y()*v2.Z()-v1.Z()*v2.Y(),
		v1.Z()*v2.X()-v1.X()*v2.Z(),
		v1.X()*v2.Y()-v1.Y()*v2.X())
}

func (t1 *Tuple) Equals(t2 *Tuple) bool {
	return equals(t1.X(), t2.X()) && equals(t1.Y(), t2.Y()) &&
		equals(t1.Z(), t2.Z()) && equals(t1.W(), t2.W())
}

func (t1 *Tuple) Transform(m *Matrix) *Tuple {
	return m.MultiplyWithTuple(t1)
}
