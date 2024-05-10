package primitives

import (
	"math"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

type Sphere struct {
	Transform *core.Matrix
	// Transform.Invert() is used for calculating every hit.
	// Since it is static we dont need to recompute it every time.
	Invert   *core.Matrix
	Material *shader.Material
}

var center *core.Tuple = core.NewPoint(0, 0, 0)

func NewSphere() *Sphere {
	s := Sphere{
		Transform: core.Identity4(),
		Invert:    core.Identity4().Invert(),
		Material:  shader.NewMaterial(),
	}
	return &s
}
func (s *Sphere) Intersect(ray *Ray) []Intersection {
	xs := make([]Intersection, 0)
	transformedRay := ray.Transform(s.Invert)
	centerToRay := transformedRay.Origin.Subtract(center)
	a := transformedRay.Direction.Dot(transformedRay.Direction)
	b := 2 * transformedRay.Direction.Dot(centerToRay)
	c := centerToRay.Dot(centerToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return xs
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	xs = append(xs, NewIntersection(t1, s))
	xs = append(xs, NewIntersection(t2, s))

	return xs
}

func (s *Sphere) SetTransform(m *core.Matrix) {
	s.Transform = m
	s.Invert = m.Invert()
}

func (s *Sphere) NormalAt(worldPoint *core.Tuple) *core.Tuple {
	objectPoint := s.Invert.MultiplyWithTuple(worldPoint)
	objectNormal := objectPoint.Subtract(center)
	worldNormal := s.Invert.Transpose().MultiplyWithTuple(objectNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()

}
