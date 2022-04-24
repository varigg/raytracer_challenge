package foundation

import (
	"fmt"
	"math"
)

const epsilon float64 = 0.00001
const vector = 0
const point = 1

type Tuple [4]float64

func (t *Tuple) X() float64 {
	return t[0]
}
func (t *Tuple) Y() float64 {
	return t[1]
}
func (t *Tuple) Z() float64 {
	return t[2]
}
func (t *Tuple) W() float64 {
	return t[3]
}

func (t *Tuple) IsPoint() bool {
	return t[3] == point
}

func (t *Tuple) IsVector() bool {
	return t[3] == vector
}

func NewTuple(x, y, z, w float64) *Tuple {
	t := Tuple{x, y, z, w}
	return &t
}

func FromColor(color *Color) *Tuple {
	return NewTuple(color[0], color[1], color[2], -1)
}
func NewPoint(x, y, z float64) *Tuple {
	t := Tuple{x, y, z, point}
	return &t
}

func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, vector}
}

func (t1 *Tuple) Plus(t2 *Tuple) *Tuple {
	return &Tuple{
		t1[0] + t2[0],
		t1[1] + t2[1],
		t1[2] + t2[2],
		t1[3] + t2[3],
	}
}

func (t1 *Tuple) Minus(t2 *Tuple) *Tuple {
	return &Tuple{
		t1[0] - t2[0],
		t1[1] - t2[1],
		t1[2] - t2[2],
		t1[3] - t2[3],
	}
}

func (t *Tuple) Negate() *Tuple {
	return &Tuple{
		0 - t[0],
		0 - t[1],
		0 - t[2],
		0 - t[3],
	}
}

func (t *Tuple) Times(s float64) *Tuple {
	return &Tuple{
		s * t[0],
		s * t[1],
		s * t[2],
		s * t[3],
	}
}

func (t *Tuple) DivideBy(s float64) *Tuple {
	return &Tuple{
		t[0] / s,
		t[1] / s,
		t[2] / s,
		t[3] / s,
	}
}

func (v *Tuple) Magnitude() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])
}

func (v *Tuple) Normalize() *Tuple {
	m := v.Magnitude()
	return NewTuple(v[0]/m, v[1]/m, v[2]/m, v[3]/m)
}

func (a *Tuple) DotProduct(b *Tuple) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

func (a *Tuple) CrossProduct(b *Tuple) (*Tuple, error) {
	if !a.IsVector() || !b.IsVector() {
		return nil, fmt.Errorf("onl[1] implemented for 3 dimensional vectors")
	}
	return NewVector(a[1]*b[2]-a[2]*b[1], a[2]*b[0]-a[0]*b[2], a[0]*b[1]-a[1]*b[0]), nil
}

func (a *Tuple) Equals(b *Tuple) bool {
	return math.Abs(a[0]-b[0]) < epsilon &&
		math.Abs(a[1]-b[1]) < epsilon &&
		math.Abs(a[2]-b[2]) < epsilon &&
		math.Abs(a[3]-b[3]) < epsilon
}
