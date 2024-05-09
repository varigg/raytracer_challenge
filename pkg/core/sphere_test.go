package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
)

func TestIntersect(t *testing.T) {
	r := core.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	s := core.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 4.0, xs[0].T)
	assert.Equal(t, 6.0, xs[1].T)
}

func TestIntersectTangent(t *testing.T) {
	r := core.NewRay(core.NewPoint(0, 1, -5), core.NewVector(0, 0, 1))
	s := core.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 5.0, xs[0].T)
	assert.Equal(t, 5.0, xs[1].T)
}

func TestIntersectMisses(t *testing.T) {
	r := core.NewRay(core.NewPoint(0, 2, -5), core.NewVector(0, 0, 1))
	s := core.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 0, len(xs))
}

func TestIntersectInside(t *testing.T) {
	r := core.NewRay(core.NewPoint(0, 0, 0), core.NewVector(0, 0, 1))
	s := core.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -1.0, xs[0].T)
	assert.Equal(t, 1.0, xs[1].T)
}

func TestIntersectBehind(t *testing.T) {
	r := core.NewRay(core.NewPoint(0, 0, 5), core.NewVector(0, 0, 1))
	s := core.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -6.0, xs[0].T)
	assert.Equal(t, -4.0, xs[1].T)

}

func TestIntersection(t *testing.T) {
	s := core.NewSphere()
	i1 := core.NewIntersection(3.5, s)
	assert.Equal(t, s, i1.Object)
	assert.Equal(t, 3.5, i1.T)
}

func TestIntersectSetsObject(t *testing.T) {
	r := core.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	s := core.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, s, xs[0].Object)
	assert.Equal(t, s, xs[1].Object)
}

func TestTransform(t *testing.T) {
	s := core.NewSphere()
	assert.Equal(t, core.Identity(4), s.Transform)
	m := core.TranslationMatrix(2, 3, 4)
	s.SetTransform(m)
	assert.Equal(t, m, s.Transform)
}

func TestIntersectScaled(t *testing.T) {
	o := core.NewPoint(0, 0, -5)
	d := core.NewVector(0, 0, 1)
	r := core.NewRay(o, d)
	s := core.NewSphere()
	s.SetTransform(core.ScalingMatrix(2, 2, 2))
	xs := s.Intersect(r)
	assert.Equal(t, float64(3), xs[0].T)
	assert.Equal(t, float64(7), xs[1].T)
}

func TestIntersectTranslated(t *testing.T) {
	o := core.NewPoint(0, 0, -5)
	d := core.NewVector(0, 0, 1)
	r := core.NewRay(o, d)
	s := core.NewSphere()
	s.SetTransform(core.TranslationMatrix(5, 0, 0))
	xs := s.Intersect(r)
	assert.Empty(t, xs)

}
