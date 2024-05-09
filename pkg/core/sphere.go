package core

import "math"

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
	xs := make([]Intersection, 0)
	transformedRay := ray.Transform(s.Transform.Invert())
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

func (s *Sphere) SetTransform(m *Matrix) {
	s.Transform = m
}
