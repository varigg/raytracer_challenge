package primitives_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/primitives"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

func TestIntersect(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	s := primitives.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 4.0, xs[0].T)
	assert.Equal(t, 6.0, xs[1].T)
}

func TestIntersectTangent(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(0, 1, -5), core.NewVector(0, 0, 1))
	s := primitives.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 5.0, xs[0].T)
	assert.Equal(t, 5.0, xs[1].T)
}

func TestIntersectMisses(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(0, 2, -5), core.NewVector(0, 0, 1))
	s := primitives.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 0, len(xs))
}

func TestIntersectInside(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(0, 0, 0), core.NewVector(0, 0, 1))
	s := primitives.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -1.0, xs[0].T)
	assert.Equal(t, 1.0, xs[1].T)
}

func TestIntersectBehind(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(0, 0, 5), core.NewVector(0, 0, 1))
	s := primitives.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -6.0, xs[0].T)
	assert.Equal(t, -4.0, xs[1].T)

}

func TestIntersection(t *testing.T) {
	s := primitives.NewSphere()
	i1 := primitives.NewIntersection(3.5, s)
	assert.Equal(t, s, i1.Object)
	assert.Equal(t, 3.5, i1.T)
}

func TestIntersectSetsObject(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	s := primitives.NewSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, s, xs[0].Object)
	assert.Equal(t, s, xs[1].Object)
}

func TestTransform(t *testing.T) {
	s := primitives.NewSphere()
	assert.Equal(t, core.Identity(4), s.Transform)
	m := core.TranslationMatrix(2, 3, 4)
	s.SetTransform(m)
	assert.Equal(t, m, s.Transform)
}

func TestIntersectScaled(t *testing.T) {
	o := core.NewPoint(0, 0, -5)
	d := core.NewVector(0, 0, 1)
	r := primitives.NewRay(o, d)
	s := primitives.NewSphere()
	s.SetTransform(core.ScalingMatrix(2, 2, 2))
	xs := s.Intersect(r)
	assert.Equal(t, float64(3), xs[0].T)
	assert.Equal(t, float64(7), xs[1].T)
}

func TestIntersectTranslated(t *testing.T) {
	o := core.NewPoint(0, 0, -5)
	d := core.NewVector(0, 0, 1)
	r := primitives.NewRay(o, d)
	s := primitives.NewSphere()
	s.SetTransform(core.TranslationMatrix(5, 0, 0))
	xs := s.Intersect(r)
	assert.Empty(t, xs)

}

func TestNormalAtXAxis(t *testing.T) {
	s := primitives.NewSphere()
	nv := s.NormalAt(core.NewPoint(1, 0, 0))
	assert.True(t, nv.Equals(core.NewVector(1, 0, 0)))
}

func TestNormalAtYAxis(t *testing.T) {
	s := primitives.NewSphere()
	nv := s.NormalAt(core.NewPoint(0, 1, 0))
	assert.True(t, nv.Equals(core.NewVector(0, 1, 0)))
}
func TestNormalAtZAxis(t *testing.T) {
	s := primitives.NewSphere()
	nv := s.NormalAt(core.NewPoint(0, 0, 1))
	assert.True(t, nv.Equals(core.NewVector(0, 0, 1)))
}

func TestNormalAtNonAxial(t *testing.T) {
	s := primitives.NewSphere()
	nv := s.NormalAt(core.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	assert.True(t, nv.Equals(core.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)))
	assert.True(t, nv.Equals(nv.Normalize()))
}

func TestNormalTranslated(t *testing.T) {
	s := primitives.NewSphere()
	s.SetTransform(core.TranslationMatrix(0, 1, 0))
	nv := s.NormalAt(core.NewPoint(0, 1.70711, -0.70711))
	assert.True(t, nv.Equals(core.NewVector(0, 0.70711, -0.70711)))
}

func TestMaterial(t *testing.T) {
	s := primitives.NewSphere()
	assert.Equal(t, shader.NewMaterial(), s.Material)
	m := shader.NewMaterial()
	m.Ambient = 1
	s.Material = m
	assert.Equal(t, m, s.Material)
}
