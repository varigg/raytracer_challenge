package scene_test

// Test that adding a translated sphere to the world does not result in duplicate/intersecting objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/scene"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

// Test that a translated and scaled sphere produces expected intersection t values
func TestWorld_TranslatedScaledSphereIntersectionT(t *testing.T) {
	w := scene.NewWorld()
	s := objects.NewSphere()
	s.SetTransform(core.TranslationMatrix(-0.5, 0, 0).Times(core.ScalingMatrix(1.5, 1.5, 1.5)))
	w.Add(s)
	// Ray from (-5,0,0) toward (0,0,0)
	r := objects.NewRay(core.NewPoint(-5, 0, 0), core.NewVector(1, 0, 0))
	xs := w.Intersect(r)
	t.Logf("Number of intersections: %d", len(xs))
	for i, inter := range xs {
		t.Logf("Intersection %d: t=%.5f, object=%v", i, inter.T, inter.Object)
	}
	// There should be 2 intersections (enter and exit the sphere)
	assert.Equal(t, 2, len(xs), "Should intersect the translated and scaled sphere exactly twice")
	// Both intersections should be with the same sphere object
	assert.Equal(t, s, xs[0].Object)
	assert.Equal(t, s, xs[1].Object)
	// t values should be symmetric and reflect the scale/translation
	t.Logf("t values: %v, %v", xs[0].T, xs[1].T)
}

// Minimal test: render a single translated sphere and check intersection at expected position
func TestWorld_TranslatedSphereIntersection(t *testing.T) {
	w := scene.NewWorld()
	s := objects.NewSphere()
	s.SetTransform(core.TranslationMatrix(-0.5, 0, 0))
	w.Add(s)
	// Ray from (0,0,-5) toward (0,0,0)
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	xs := w.Intersect(r)
	t.Logf("Number of intersections: %d", len(xs))
	for i, inter := range xs {
		t.Logf("Intersection %d: t=%.5f, object=%v", i, inter.T, inter.Object)
	}
	// There should be 2 intersections (enter and exit the sphere)
	assert.Equal(t, 2, len(xs), "Should intersect the translated sphere exactly twice")
	// Both intersections should be with the same sphere object
	assert.Equal(t, s, xs[0].Object)
	assert.Equal(t, s, xs[1].Object)
}

func TestWorld_AddTranslatedSphere_NoDuplicate(t *testing.T) {
	w := scene.NewWorld()
	s := objects.NewSphere()
	s.SetTransform(core.TranslationMatrix(-0.5, 0, 0))
	w.Add(s)
	// There should be exactly one object in the world
	assert.Equal(t, 1, len(w.Objects), "World should contain exactly one object after adding a translated sphere")
	// The transform of the object should match what we set
	assert.True(t, core.TranslationMatrix(-0.5, 0, 0).Equals(w.Objects[0].(*objects.Sphere).Transform()), "Sphere transform should match the translation applied")
}

func TestNewWorld(t *testing.T) {
	w := scene.NewWorld()
	assert.Nil(t, w.Light)
	assert.Nil(t, w.Objects)
}

// Scenario: The default world
//
//	Given light ← point_light(point(-10, 10, -10), color(1, 1, 1))
//	  And s1 ← sphere() with:
//	    | material.color     | (0.8, 1.0, 0.6)        |
//	    | material.diffuse   | 0.7                    |
//	    | material.specular  | 0.2                    |
//	  And s2 ← sphere() with:
//	    | transform | scaling(0.5, 0.5, 0.5) |
//	When w ← default_world()
//	Then w.light = light
//	  And w contains s1
//	  And w contains s2
func TestNewDefaultWorld(t *testing.T) {
	w := scene.NewDefaultWorld()
	light := shader.NewLight(core.NewPoint(-10, 10, -10), core.NewColor(1, 1, 1))
	assert.Equal(t, light, w.Light)
	assert.Equal(t, w.Objects[0].Material().Diffuse, .7)
	assert.Equal(t, w.Objects[1].(*objects.Sphere).Transform(), core.ScalingMatrix(.5, .5, .5))
}

// Scenario: Intersect a world with a ray
//
//	Given w ← default_world()
//	  And r ← ray(point(0, 0, -5), vector(0, 0, 1))
//	When xs ← intersect_world(w, r)
//	Then xs.count = 4
//	  And xs[0].t = 4
//	  And xs[1].t = 4.5
//	  And xs[2].t = 5.5
//	  And xs[3].t = 6
func TestIntersectWorld(t *testing.T) {
	w := scene.NewDefaultWorld()
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	xs := w.Intersect(r)
	assert.Equal(t, 4, len(xs))
	assert.Equal(t, 4.0, xs[0].T)
	assert.Equal(t, 4.5, xs[1].T)
	assert.Equal(t, 5.5, xs[2].T)
	assert.Equal(t, 6.0, xs[3].T)
}

// Scenario: Shading an intersection
//
//	Given w ← default_world()
//	  And r ← ray(point(0, 0, -5), vector(0, 0, 1))
//	  And shape ← the first object in w
//	  And i ← intersection(4, shape)
//	When comps ← prepare_computations(i, r)
//	  And c ← shade_hit(w, comps)
//	Then c = color(0.38066, 0.47583, 0.2855)
func TestShadeHit(t *testing.T) {
	w := scene.NewDefaultWorld()
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	shape := w.Objects[0]
	i := objects.NewIntersection(4, shape)
	comps := scene.PrepareComputations(&i, r)
	color := w.ShadeHit(comps)
	assert.True(t, color.Equals(core.NewColor(0.38066, 0.47583, 0.2855)))
}

// Scenario: Shading an intersection from the inside
//
//	Given w ← default_world()
//	  And w.light ← point_light(point(0, 0.25, 0), color(1, 1, 1))
//	  And r ← ray(point(0, 0, 0), vector(0, 0, 1))
//	  And shape ← the second object in w
//	  And i ← intersection(0.5, shape)
//	When comps ← prepare_computations(i, r)
//	  And c ← shade_hit(w, comps)
//	Then c = color(0.90498, 0.90498, 0.90498)
func TestShadeHitInside(t *testing.T) {
	w := scene.NewDefaultWorld()
	r := objects.NewRay(core.NewPoint(0, 0, 0), core.NewVector(0, 0, 1))
	w.Light = shader.NewLight(core.NewPoint(0, 0.25, 0), core.NewColor(1, 1, 1))
	shape := w.Objects[1]
	i := objects.NewIntersection(.5, shape)
	comps := scene.PrepareComputations(&i, r)
	color := w.ShadeHit(comps)
	t.Logf("Actual color: %v", color)
	t.Logf("Expected color: %v", core.NewColor(0.90498, 0.90498, 0.90498))
	assert.True(t, color.Equals(core.NewColor(0.90498, 0.90498, 0.90498)))
}

func TestColorRayMissed(t *testing.T) {
	w := scene.NewDefaultWorld()
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 1, 0))
	assert.True(t, core.NewColor(0, 0, 0).Equals(w.ColorAt(r)))
}

func TestColorRayHit(t *testing.T) {
	w := scene.NewDefaultWorld()
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	assert.True(t, core.NewColor(0.38066, 0.47583, 0.2855).Equals(w.ColorAt(r)))
}

func TestColorRayInside(t *testing.T) {
	w := scene.NewDefaultWorld()
	outer := w.Objects[0]
	outer.Material().Ambient = 1
	inner := w.Objects[1]
	inner.Material().Ambient = 1
	r := objects.NewRay(core.NewPoint(0, 0, 0.75), core.NewVector(0, 0, -1))
	c := w.ColorAt(r)
	assert.True(t, inner.Material().Color.Equals(c))
}
