package scene_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/scene"
)

// Test that ColorAt returns the color of the object actually hit by the ray
func TestColorAt_HitsCorrectObject(t *testing.T) {
	w := scene.NewDefaultWorld()
	outer := w.Objects[0]
	inner := w.Objects[1]
	outer.GetMaterial().Ambient = 1
	inner.GetMaterial().Ambient = 1

	r := objects.NewRay(core.NewPoint(0, 0, 0.75), core.NewVector(0, 0, -1))
	color := w.ColorAt(r)
	// The ray hits the inner sphere first, so expect its color
	assert.True(t, inner.GetMaterial().Color.Equals(color), "ColorAt should return the color of the inner object hit by the ray")
}

// Test that when ambient is 1, the result is exactly the material color
func TestShadeHit_AmbientOnly(t *testing.T) {
	w := scene.NewDefaultWorld()
	shape := w.Objects[0]
	shape.GetMaterial().Ambient = 1
	shape.GetMaterial().Diffuse = 0
	shape.GetMaterial().Specular = 0
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	i := objects.NewIntersection(4, shape)
	comps := scene.PrepareComputations(&i, r)
	color := w.ShadeHit(comps)
	assert.True(t, shape.GetMaterial().Color.Equals(color), "ShadeHit should return the material color when ambient=1 and others=0")
}

// Test that the intersection order is correct when the ray starts inside an object
func TestIntersectionOrder_RayStartsInside(t *testing.T) {
	w := scene.NewDefaultWorld()
	t.Logf("Outer sphere transform: %v", w.Objects[0].GetTransform())
	t.Logf("Inner sphere transform: %v", w.Objects[1].GetTransform())
	// Start the ray inside the inner sphere (radius 0.5, centered at origin)
	r := objects.NewRay(core.NewPoint(0, 0, 0.1), core.NewVector(0, 0, 1))
	xs := w.Intersect(r)
	t.Logf("Number of intersections: %d", len(xs))
	for i, inter := range xs {
		which := "outer"
		if inter.Object == w.Objects[1] {
			which = "inner"
		}
		t.Logf("Intersection %d: t=%.5f, object=%s", i, inter.T, which)
	}
	assert.Greater(t, len(xs), 0, "There should be intersections")
	// The first positive intersection (the hit) should be with the inner sphere
	var hitObj interface{} = nil
	for _, inter := range xs {
		if inter.T >= 0 {
			hitObj = inter.Object
			break
		}
	}
	assert.Equal(t, w.Objects[1], hitObj, "The first positive intersection (the hit) should be with the inner sphere")
}
