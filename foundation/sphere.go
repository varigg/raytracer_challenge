package foundation

import (
	"math"
)

type Sphere struct {
	Transform *Matrix
}

var center *Tuple = NewPoint(0, 0, 0)

func NewSphere() *Sphere {
	s := Sphere{
		Transform: Identity4(),
	}
	return &s
}

func (s *Sphere) Intersect(ray *Ray) []Intersection {
	transformedRay := ray.Transform(s.Transform.Invert())
	centerToRay := transformedRay.Origin.Minus(center)
	a := transformedRay.Direction.DotProduct(transformedRay.Direction)
	b := 2 * transformedRay.Direction.DotProduct(centerToRay)
	c := centerToRay.DotProduct(centerToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	i1 := NewIntersection(t1, s)
	i2 := NewIntersection(t2, s)
	return []Intersection{i1, i2}
}

func (s *Sphere) SetTransform(m *Matrix) {
	s.Transform = m
}
