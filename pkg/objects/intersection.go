package objects

type Intersection struct {
	T      float64
	Object Object
}

func NewIntersection(t float64, object Object) Intersection {
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
