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
}

// 	Divide(float64) Tuple
// 	Magnitude() float64
// 	Normalize() Tuple
// 	Dot(Tuple) float64
// 	Cross(Tuple) Tuple
// }

type tuple struct {
	x float64
	y float64
	z float64
	w float64
}

type Point struct {
	Tuple
}
type Vector struct {
	Tuple
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

func NewPoint(a, b, c float64) *Point {
	p := &Point{
		NewTuple(a, b, c, 1.0),
		// t: tuple{
		// 	x: a,
		// 	y: b,
		// 	z: c,
		// 	w: 1.0,
		// },
	}
	return p
}

func NewVector(a, b, c float64) *Vector {
	// t := tuple{
	// 	x: a,
	// 	y: b,
	// 	z: c,
	// 	w: 0.0,
	// }
	v := &Vector{
		NewTuple(a, b, c, 0.0),
	}
	return v
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

func NewVectorFromString(v string) *Vector {
	x, y, z := stringToCoordinates(v)
	return NewVector(x, y, z)
}

func NewPointFromString(v string) *Point {
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

func (p *Point) Add(t2 Tuple) Tuple {
	t := &tuple{
		x: p.X() + t2.X(),
		y: p.Y() + t2.Y(),
		z: p.Z() + t2.Z(),
		w: p.W() + t2.W(),
	}
	return &Point{
		t,
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

func (t1 *Point) Subtract(t2 Tuple) Tuple {
	t := &tuple{
		x: t1.X() - t2.X(),
		y: t1.Y() - t2.Y(),
		z: t1.Z() - t2.Z(),
		w: t1.W() - t2.W(),
	}
	if t.w == 0.0 {
		return NewVector(t.x, t.y, t.z)
	}
	if t.w == 1.0 {
		return NewPoint(t.x, t.y, t.z)
	}
	return t
}

func (t1 *Vector) Subtract(t2 Tuple) Tuple {

	return NewVector(t1.X()-t2.X(), t1.Y()-t2.Y(), t1.Z()-t2.Z())
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

func (v *Vector) Multiply(s float64) Tuple {
	return NewVector(v.X()*s, v.Y()*s, v.Z()*s)
}

func (t *tuple) Divide(s float64) Tuple {
	return &tuple{
		x: t.x / s,
		y: t.y / s,
		z: t.z / s,
		w: t.w / s,
	}
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z() + v.W()*v.W())
}

func (v *Vector) Normalize() *Vector {
	t := NewTuple(v.X()/v.Magnitude(),
		v.Y()/v.Magnitude(),
		v.Z()/v.Magnitude(),
		v.W()/v.Magnitude(),
	)
	return &Vector{
		t,
	}
}

func (t1 *Vector) Dot(t2 *Vector) float64 {
	return t1.X()*t2.X() + t1.Y()*t2.Y() +
		t1.Z()*t2.Z() + t1.W()*t2.W()
}

func (v1 *Vector) Cross(v2 *Vector) *Vector {
	// This makes only sense for vectors
	return NewVector(v1.Y()*v2.Z()-v1.Z()*v2.Y(),
		v1.Z()*v2.X()-v1.X()*v2.Z(),
		v1.X()*v2.Y()-v1.Y()*v2.X())
}
