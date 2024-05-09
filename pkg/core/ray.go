package core

type Ray struct {
	Origin        *Tuple
	Direction     *Tuple
	Intersections []Intersection
	hit           *Intersection
}

type Intersecter interface {
	Intersect(*Ray) []Intersection
}

type Intersection struct {
	T      float64
	Object Intersecter
}

func NewRay(origin, direction *Tuple) *Ray {
	r := Ray{
		Origin:    origin,
		Direction: direction,
	}
	return &r
}

func (r *Ray) Position(t float64) *Tuple {
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

func (r *Ray) Transform(m *Matrix) *Ray {
	newRay := Ray{
		Origin:    m.MultiplyWithTuple(r.Origin),
		Direction: m.MultiplyWithTuple(r.Direction),
	}
	return &newRay
}
