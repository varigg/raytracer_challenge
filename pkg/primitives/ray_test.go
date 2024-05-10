package primitives_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/primitives"
)

func TestRay(t *testing.T) {

	o := core.NewPoint(1, 2, 3)
	d := core.NewVector(4, 5, 6)
	r := primitives.NewRay(o, d)
	assert.Equal(t, o, r.Origin)
	assert.Equal(t, d, r.Direction)
}

func TestPosition(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(2, 3, 4), core.NewVector(1, 0, 0))
	assert.Equal(t, core.NewPoint(2, 3, 4), r.Position(0))
	assert.Equal(t, core.NewPoint(3, 3, 4), r.Position(1))
	assert.Equal(t, core.NewPoint(1, 3, 4), r.Position(-1))
	assert.Equal(t, core.NewPoint(4.5, 3, 4), r.Position(2.5))
}

func TestIntersections(t *testing.T) {
	s := primitives.NewSphere()
	i := primitives.NewIntersection(3.5, s)
	assert.Equal(t, 3.5, i.T)
	assert.Equal(t, s, i.Object)

}

func TestHitPositive(t *testing.T) {
	s := primitives.NewSphere()
	r := primitives.NewRay(core.NewPoint(0, 0, 1), core.NewVector(0, 0, 1))
	r.AddIntersections(primitives.NewIntersection(1, s))
	r.AddIntersections(primitives.NewIntersection(2, s))
	assert.Equal(t, primitives.NewIntersection(1, s), (*r.Hit()))
}

func TestHitNegative(t *testing.T) {
	s := primitives.NewSphere()
	r := primitives.NewRay(core.NewPoint(0, 0, 1), core.NewVector(0, 0, 1))
	r.AddIntersections(primitives.NewIntersection(-1, s))
	r.AddIntersections(primitives.NewIntersection(-2, s))
	assert.Nil(t, r.Hit())
}

func TestHitMixed(t *testing.T) {
	s := primitives.NewSphere()
	r := primitives.NewRay(core.NewPoint(0, 0, 1), core.NewVector(0, 0, 1))
	r.AddIntersections(primitives.NewIntersection(1, s))
	r.AddIntersections(primitives.NewIntersection(-1, s))
	assert.Equal(t, primitives.NewIntersection(1, s), (*r.Hit()))
}

func TestHitMixedMultipePositive(t *testing.T) {
	s := primitives.NewSphere()
	r := primitives.NewRay(core.NewPoint(0, 0, 1), core.NewVector(0, 0, 1))
	r.AddIntersections(primitives.NewIntersection(5, s))
	r.AddIntersections(primitives.NewIntersection(7, s))
	r.AddIntersections(primitives.NewIntersection(-3, s))
	r.AddIntersections(primitives.NewIntersection(2, s))
	assert.Equal(t, primitives.NewIntersection(2, s), (*r.Hit()))
}

func TestRayTranslation(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(1, 2, 3), core.NewVector(0, 1, 0))
	m := core.TranslationMatrix(3, 4, 5)
	r2 := r.Transform(m)
	assert.Equal(t, core.NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, core.NewVector(0, 1, 0), r2.Direction)

}

func TestRayScaling(t *testing.T) {
	r := primitives.NewRay(core.NewPoint(1, 2, 3), core.NewVector(0, 1, 0))
	m := core.ScalingMatrix(2, 3, 4)
	r2 := r.Transform(m)
	assert.Equal(t, core.NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, core.NewVector(0, 3, 0), r2.Direction)

}
