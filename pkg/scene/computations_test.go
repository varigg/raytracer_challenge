package scene_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/scene"
)

// Scenario: Precomputing the state of an intersection
//
//	Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
//	  And shape ← sphere()
//	  And i ← intersection(4, shape)
//	When comps ← prepare_computations(i, r)
//	Then comps.t = i.t
//	  And comps.object = i.object
//	  And comps.point = point(0, 0, -1)
//	  And comps.eyev = vector(0, 0, -1)
//	  And comps.normalv = vector(0, 0, -1)
//
// Scenario: The hit, when an intersection occurs on the outside
//
//	Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
//	  And shape ← sphere()
//	  And i ← intersection(4, shape)
//	When comps ← prepare_computations(i, r)
//	Then comps.inside = false
func TestPreComputeIntersection(t *testing.T) {
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	shape := objects.NewSphere()
	i := objects.NewIntersection(4, shape)
	comps := scene.PrepareComputations(&i, r)
	assert.Equal(t, shape, comps.Object)
	assert.Equal(t, i.T, comps.T)
	assert.Equal(t, core.NewPoint(0, 0, -1), comps.Point)
	assert.Equal(t, core.NewVector(0, 0, -1), comps.EyeV)
	assert.Equal(t, core.NewVector(0, 0, -1), comps.NormalV)
	assert.Equal(t, false, comps.Inside)
}

// Scenario: The hit, when an intersection occurs on the inside
//
//	Given r ← ray(point(0, 0, 0), vector(0, 0, 1))
//	  And shape ← sphere()
//	  And i ← intersection(1, shape)
//	When comps ← prepare_computations(i, r)
//	Then comps.point = point(0, 0, 1)
//	  And comps.eyev = vector(0, 0, -1)
//	  And comps.inside = true
//	  And comps.normalv = vector(0, 0, -1)
func TestPreComputeIntersectionInside(t *testing.T) {
	r := objects.NewRay(core.NewPoint(0, 0, 0), core.NewVector(0, 0, 1))
	shape := objects.NewSphere()
	i := objects.NewIntersection(1, shape)
	comps := scene.PrepareComputations(&i, r)
	assert.Equal(t, shape, comps.Object)
	assert.Equal(t, i.T, comps.T)
	assert.Equal(t, core.NewPoint(0, 0, 1), comps.Point)
	assert.Equal(t, core.NewVector(0, 0, -1), comps.EyeV)
	assert.Equal(t, core.NewVector(0, 0, -1), comps.NormalV)
	assert.Equal(t, true, comps.Inside)
}
