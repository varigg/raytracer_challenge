package primitives

import "github.com/varigg/raytracer-challenge/pkg/core"

type Ray struct {
	Origin        *core.Tuple
	Direction     *core.Tuple
	Intersections []Intersection
	hit           *Intersection
}

type Intersecter interface {
	Intersect(*Ray) []Intersection
	NormalAt(*core.Tuple) *core.Tuple
}

type Intersection struct {
	T      float64
	Object Intersecter
}

func NewRay(origin, direction *core.Tuple) *Ray {
	r := Ray{
		Origin:    origin,
		Direction: direction,
	}
	return &r
}

func (r *Ray) Position(t float64) *core.Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func NewIntersection(t float64, object Intersecter) Intersection {
	i := Intersection{
		T:      t,
		Object: object,
	}
	return i
}

func (r *Ray) AddIntersections(xs ...Intersection) {
	for _, i := range xs {
		if i.T > 0 && (r.hit == nil || i.T < r.hit.T) {
			r.hit = &i
		}
		r.Intersections = append(r.Intersections, i)
	}
}

func (r *Ray) Hit() *Intersection {
	return r.hit
}

func (r *Ray) Transform(m *core.Matrix) *Ray {
	newRay := Ray{
		Origin:    m.MultiplyWithTuple(r.Origin),
		Direction: m.MultiplyWithTuple(r.Direction),
	}
	return &newRay
}
