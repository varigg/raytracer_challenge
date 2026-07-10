package objects

import (
	"math"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

type Sphere struct {
	transform core.Matrix
	// Transform.Invert() is used for calculating every hit.
	// Since it is static we dont need to recompute it every time.
	invert   core.Matrix
	Material *shader.Material
}

var center = core.NewPoint(0, 0, 0)

func NewSphere() *Sphere {
	idInv := core.Identity(4).Invert()
	s := Sphere{
		transform: core.Identity(4),
		invert:    idInv,
		Material:  shader.NewMaterial(),
	}
	return &s
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

func (s *Sphere) SetTransform(m core.Matrix) error {
	s.transform = m
	s.invert = m.Invert()
	return nil
}

func (s *Sphere) GetTransform() core.Matrix {
	return s.transform
}

func (s *Sphere) GetInverseTransform() core.Matrix {
	return s.invert
}

func (s *Sphere) NormalAt(worldPoint core.Tuple) core.Tuple {
	objectPoint := s.invert.MultiplyWithTuple(worldPoint)
	objectNormal := objectPoint.Subtract(center)
	worldNormal := s.invert.Transpose().MultiplyWithTuple(objectNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()

}

func (s *Sphere) GetMaterial() *shader.Material {
	return s.Material
}
