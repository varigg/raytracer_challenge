package foundation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytacer-challenge/foundation"
)

func TestIntersectCenter(t *testing.T) {
	o := foundation.NewPoint(0, 0, -5)
	d := foundation.NewVector(0, 0, 1)
	r := foundation.NewRay(o, d)
	s := foundation.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 4.0, xs[0].T)
	assert.Equal(t, 6.0, xs[1].T)
}

func TestIntersectTangent(t *testing.T) {
	o := foundation.NewPoint(0, 1, -5)
	d := foundation.NewVector(0, 0, 1)
	r := foundation.NewRay(o, d)
	s := foundation.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 5.0, xs[0].T)
	assert.Equal(t, 5.0, xs[1].T)
}

func TestIntersectMisses(t *testing.T) {
	o := foundation.NewPoint(0, 2, -5)
	d := foundation.NewVector(0, 0, 1)
	r := foundation.NewRay(o, d)
	s := foundation.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 0, len(xs))
}

func TestIntersectStartInside(t *testing.T) {
	o := foundation.NewPoint(0, 0, 0)
	d := foundation.NewVector(0, 0, 1)
	r := foundation.NewRay(o, d)
	s := foundation.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -1.0, xs[0].T)
	assert.Equal(t, 1.0, xs[1].T)
}

func TestIntersectBehind(t *testing.T) {
	o := foundation.NewPoint(0, 0, 5)
	d := foundation.NewVector(0, 0, 1)
	r := foundation.NewRay(o, d)
	s := foundation.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -6.0, xs[0].T)
	assert.Equal(t, -4.0, xs[1].T)
	assert.Equal(t, s, xs[1].Object)
}

func TestTransform(t *testing.T) {
	s := foundation.NewSphere()
	assert.Equal(t, foundation.Identity(4), s.Transform)
	m := foundation.TranslationMatrix(2, 3, 4)
	s.SetTransform(m)
	assert.Equal(t, m, s.Transform)
}

func TestIntersectScaled(t *testing.T) {
	o := foundation.NewPoint(0, 0, -5)
	d := foundation.NewVector(0, 0, 1)
	r := foundation.NewRay(o, d)
	s := foundation.NewSphere()
	s.SetTransform(foundation.ScalingMatrix(2, 2, 2))
	xs := s.Intersect(r)
	assert.Equal(t, float64(3), xs[0].T)
	assert.Equal(t, float64(7), xs[1].T)
}

func TestIntersectTranslated(t *testing.T) {
	o := foundation.NewPoint(0, 0, -5)
	d := foundation.NewVector(0, 0, 1)
	r := foundation.NewRay(o, d)
	s := foundation.NewSphere()
	s.SetTransform(foundation.TranslationMatrix(5, 0, 0))
	xs := s.Intersect(r)
	assert.Empty(t, xs)

}
