package foundation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytacer-challenge/foundation"
)

func TestRay(t *testing.T) {
	o := foundation.NewPoint(1, 2, 3)
	d := foundation.NewVector(4, 5, 6)
	r := foundation.NewRay(o, d)
	assert.Equal(t, o, r.Origin)
	assert.Equal(t, d, r.Direction)
}

func TestPosition(t *testing.T) {
	r := foundation.NewRay(foundation.NewPoint(2, 3, 4), foundation.NewVector(1, 0, 0))
	assert.Equal(t, foundation.NewPoint(2, 3, 4), r.Position(0))
	assert.Equal(t, foundation.NewPoint(3, 3, 4), r.Position(1))
	assert.Equal(t, foundation.NewPoint(1, 3, 4), r.Position(-1))
	assert.Equal(t, foundation.NewPoint(4.5, 3, 4), r.Position(2.5))
}

func TestIntersections(t *testing.T) {
	s := foundation.NewSphere()
	i := foundation.NewIntersection(3.5, s)
	assert.Equal(t, 3.5, i.T)
	assert.Equal(t, s, i.Object)

}

func TestHitPositive(t *testing.T) {
	s := foundation.NewSphere()
	r := foundation.NewRay(foundation.NewPoint(0, 0, 1), foundation.NewVector(0, 0, 1))
	r.AddIntersections(foundation.NewIntersection(1, s))
	r.AddIntersections(foundation.NewIntersection(2, s))
	assert.Equal(t, foundation.NewIntersection(1, s), (*r.Hit()))
}

func TestHitNegative(t *testing.T) {
	s := foundation.NewSphere()
	r := foundation.NewRay(foundation.NewPoint(0, 0, 1), foundation.NewVector(0, 0, 1))
	r.AddIntersections(foundation.NewIntersection(-1, s))
	r.AddIntersections(foundation.NewIntersection(-2, s))
	assert.Nil(t, r.Hit())
}

func TestHitMixed(t *testing.T) {
	s := foundation.NewSphere()
	r := foundation.NewRay(foundation.NewPoint(0, 0, 1), foundation.NewVector(0, 0, 1))
	r.AddIntersections(foundation.NewIntersection(1, s))
	r.AddIntersections(foundation.NewIntersection(-1, s))
	assert.Equal(t, foundation.NewIntersection(1, s), (*r.Hit()))
}

func TestHitMixedMultipePositive(t *testing.T) {
	s := foundation.NewSphere()
	r := foundation.NewRay(foundation.NewPoint(0, 0, 1), foundation.NewVector(0, 0, 1))
	r.AddIntersections(foundation.NewIntersection(5, s))
	r.AddIntersections(foundation.NewIntersection(7, s))
	r.AddIntersections(foundation.NewIntersection(-3, s))
	r.AddIntersections(foundation.NewIntersection(2, s))
	assert.Equal(t, foundation.NewIntersection(2, s), (*r.Hit()))
}

func TestRayTranslation(t *testing.T) {
	r := foundation.NewRay(foundation.NewPoint(1, 2, 3), foundation.NewVector(0, 1, 0))
	m := foundation.TranslationMatrix(3, 4, 5)
	r2 := r.Transform(m)
	assert.Equal(t, foundation.NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, foundation.NewVector(0, 1, 0), r2.Direction)

}

func TestRayScaling(t *testing.T) {
	r := foundation.NewRay(foundation.NewPoint(1, 2, 3), foundation.NewVector(0, 1, 0))
	m := foundation.ScalingMatrix(2, 3, 4)
	r2 := r.Transform(m)
	assert.Equal(t, foundation.NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, foundation.NewVector(0, 3, 0), r2.Direction)

}
