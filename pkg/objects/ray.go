package objects

import "github.com/varigg/raytracer-challenge/pkg/core"

type Ray struct {
	Origin    core.Tuple
	Direction core.Tuple
}

func NewRay(origin, direction core.Tuple) *Ray {
	r := Ray{
		Origin:    origin,
		Direction: direction,
	}
	return &r
}

func (r *Ray) Position(t float64) core.Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r *Ray) Transform(m *core.Matrix) *Ray {
	newRay := Ray{
		Origin:    m.MultiplyWithTuple(r.Origin),
		Direction: m.MultiplyWithTuple(r.Direction),
	}
	return &newRay
}
