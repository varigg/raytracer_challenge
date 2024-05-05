package raytracer

import (
	"math"
	"strconv"
	"strings"
)

type Tuple interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
	IsVector() bool
	IsPoint() bool
	Add(Tuple) Tuple
	Subtract(Tuple) Tuple
	Negate() Tuple
	Multiply(float64) Tuple
	Divide(float64) Tuple
	Magnitude() float64
	Normalize() Tuple
	Dot(Tuple) float64
	Cross(Tuple) Tuple
}

type tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func NewTuple(a, b, c, w float64) Tuple {
	t := &tuple{
		x: a,
		y: b,
		z: c,
		w: w,
	}
	return t

}

func NewPoint(a, b, c float64) Tuple {
	t := &tuple{
		x: a,
		y: b,
		z: c,
		w: 1.0,
	}
	return t
}

func NewVector(a, b, c float64) Tuple {
	t := &tuple{
		x: a,
		y: b,
		z: c,
		w: 0.0,
	}
	return t
}

func (t *tuple) X() float64 {
	return t.x
}

func (t *tuple) Y() float64 {
	return t.y
}

func (t *tuple) Z() float64 {
	return t.z
}

func (t *tuple) W() float64 {
	return t.w
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

func (t *tuple) IsVector() bool {
	return t.w == 0.0
}

func (t *tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t1 *tuple) Add(t2 Tuple) Tuple {
	return &tuple{
		x: t1.x + t2.X(),
		y: t1.y + t2.Y(),
		z: t1.z + t2.Z(),
		w: t1.w + t2.W(),
	}
}

func (t1 *tuple) Subtract(t2 Tuple) Tuple {
	return &tuple{
		x: t1.x - t2.X(),
		y: t1.y - t2.Y(),
		z: t1.z - t2.Z(),
		w: t1.w - t2.W(),
	}
}

func (t *tuple) Negate() Tuple {
	return &tuple{
		x: 0 - t.x,
		y: 0 - t.y,
		z: 0 - t.z,
		w: 0 - t.w,
	}
}

func (t *tuple) Multiply(s float64) Tuple {
	return &tuple{
		x: t.x * s,
		y: t.y * s,
		z: t.z * s,
		w: t.w * s,
	}
}

func (t *tuple) Divide(s float64) Tuple {
	return &tuple{
		x: t.x / s,
		y: t.y / s,
		z: t.z / s,
		w: t.w / s,
	}
}

func (t *tuple) Magnitude() float64 {

	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z + t.w*t.w)
}

func (t *tuple) Normalize() Tuple {
	return &tuple{
		x: t.x / t.Magnitude(),
		y: t.y / t.Magnitude(),
		z: t.z / t.Magnitude(),
		w: t.w / t.Magnitude(),
	}
}

func (t1 *tuple) Dot(t2 Tuple) float64 {
	return t1.x*t2.X() + t1.y*t2.Y() +
		t1.z*t2.Z() + t1.w*t2.W()
}

func (v1 *tuple) Cross(v2 Tuple) Tuple {
	// This makes only sense for vectors
	return NewVector(v1.y*v2.(*tuple).z-v1.z*v2.(*tuple).y,
		v1.z*v2.(*tuple).x-v1.x*v2.(*tuple).z,
		v1.x*v2.(*tuple).y-v1.y*v2.(*tuple).x)
}
