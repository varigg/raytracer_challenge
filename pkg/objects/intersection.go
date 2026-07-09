package objects

import (
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

type Intersecter interface {
	Intersect(*Ray) []Intersection
	NormalAt(core.Tuple) core.Tuple
	GetMaterial() *shader.Material
}

type Intersection struct {
	T      float64
	Object Intersecter
}

func NewIntersection(t float64, object Intersecter) Intersection {
	i := Intersection{
		T:      t,
		Object: object,
	}
	return i
}

func Hit(intersections []Intersection) *Intersection {
	var hit *Intersection
	for i, intersection := range intersections {
		if intersection.T >= 0 {
			if hit == nil || intersection.T < hit.T {
				hit = &intersections[i]
			}
		}
	}
	return hit
}
