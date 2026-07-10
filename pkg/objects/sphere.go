package objects

import (
	"math"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

type Sphere struct {
	transform core.Matrix
	// invert caches transform.Invert(); every intersection needs it.
	invert   core.Matrix
	material *shader.Material
}

var center = core.NewPoint(0, 0, 0)

func NewSphere() *Sphere {
	return &Sphere{
		transform: core.Identity(4),
		invert:    core.Identity(4).Invert(),
		material:  shader.NewMaterial(),
	}
}
func (s *Sphere) Intersect(ray *Ray) []Intersection {
	xs := make([]Intersection, 0)
	transformedRay := ray.Transform(s.invert)
	centerToRay := transformedRay.Origin.Subtract(center)
	a := transformedRay.Direction.Dot(transformedRay.Direction)
	b := 2 * transformedRay.Direction.Dot(centerToRay)
	c := centerToRay.Dot(centerToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return nil
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	xs = append(xs, NewIntersection(t1, s))
	xs = append(xs, NewIntersection(t2, s))

	return xs
}

func (s *Sphere) SetTransform(m core.Matrix) {
	s.transform = m
	s.invert = m.Invert()
}

func (s *Sphere) Transform() core.Matrix         { return s.transform }
func (s *Sphere) InverseTransform() core.Matrix  { return s.invert }
func (s *Sphere) Material() *shader.Material     { return s.material }
func (s *Sphere) SetMaterial(m *shader.Material) { s.material = m }

func (s *Sphere) NormalAt(worldPoint core.Tuple) core.Tuple {
	objectPoint := s.invert.MultiplyWithTuple(worldPoint)
	objectNormal := objectPoint.Subtract(center)
	worldNormal := s.invert.Transpose().MultiplyWithTuple(objectNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()

}
