package objects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
)

func TestRay(t *testing.T) {

	o := core.NewPoint(1, 2, 3)
	d := core.NewVector(4, 5, 6)
	r := objects.NewRay(o, d)
	assert.Equal(t, o, r.Origin)
	assert.Equal(t, d, r.Direction)
}

func TestPosition(t *testing.T) {
	r := objects.NewRay(core.NewPoint(2, 3, 4), core.NewVector(1, 0, 0))
	assert.Equal(t, core.NewPoint(2, 3, 4), r.Position(0))
	assert.Equal(t, core.NewPoint(3, 3, 4), r.Position(1))
	assert.Equal(t, core.NewPoint(1, 3, 4), r.Position(-1))
	assert.Equal(t, core.NewPoint(4.5, 3, 4), r.Position(2.5))
}

func TestIntersections(t *testing.T) {
	s := objects.NewSphere()
	i := objects.NewIntersection(3.5, s)
	assert.Equal(t, 3.5, i.T)
	assert.Equal(t, s, i.Object)

}

func TestHitPositive(t *testing.T) {
	s := objects.NewSphere()
	i1 := objects.NewIntersection(1, s)
	i2 := objects.NewIntersection(2, s)
	xs := []objects.Intersection{i1, i2}
	assert.Equal(t, i1, (*objects.Hit(xs)))
}

func TestHitNegative(t *testing.T) {
	s := objects.NewSphere()
	i1 := objects.NewIntersection(-2, s)
	i2 := objects.NewIntersection(-1, s)
	xs := []objects.Intersection{i1, i2}
	assert.Nil(t, objects.Hit(xs))
}

func TestHitMixed(t *testing.T) {
	s := objects.NewSphere()
	i1 := objects.NewIntersection(-1, s)
	i2 := objects.NewIntersection(1, s)
	xs := []objects.Intersection{i1, i2}
	assert.Equal(t, i2, (*objects.Hit(xs)))
}

func TestHitMixedMultipePositive(t *testing.T) {
	s := objects.NewSphere()
	i1 := objects.NewIntersection(5, s)
	i2 := objects.NewIntersection(7, s)
	i3 := objects.NewIntersection(-3, s)
	i4 := objects.NewIntersection(2, s)
	xs := []objects.Intersection{i1, i2, i3, i4}
	assert.Equal(t, i4, (*objects.Hit(xs)))
}

func TestRayTranslation(t *testing.T) {
	r := objects.NewRay(core.NewPoint(1, 2, 3), core.NewVector(0, 1, 0))
	m := core.TranslationMatrix(3, 4, 5)
	r2 := r.Transform(m)
	assert.Equal(t, core.NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, core.NewVector(0, 1, 0), r2.Direction)

}

func TestRayScaling(t *testing.T) {
	r := objects.NewRay(core.NewPoint(1, 2, 3), core.NewVector(0, 1, 0))
	m := core.ScalingMatrix(2, 3, 4)
	r2 := r.Transform(m)
	assert.Equal(t, core.NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, core.NewVector(0, 3, 0), r2.Direction)

}
