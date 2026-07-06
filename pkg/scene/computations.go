package scene

import (
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
)

type Computations struct {
	T      float64
	Object objects.Intersecter
	Point  *core.Tuple
	EyeV   *core.Tuple
	NormalV *core.Tuple
	Inside bool
}

func PrepareComputations(i *objects.Intersection, ray *objects.Ray) *Computations {
	c := &Computations{}
	c.T = i.T
	c.Object = i.Object
	c.Point = ray.Position(i.T)
	c.EyeV = ray.Direction.Negate()
	c.NormalV = c.Object.NormalAt(c.Point)
	if c.NormalV.Dot(c.EyeV) < 0 {
		c.Inside = true
		c.NormalV = c.NormalV.Negate()
	}
	return c
}
