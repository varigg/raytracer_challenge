package objects_test

import (
	"testing"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
)

var (
	sinkIntersections []objects.Intersection
	sinkNormal        core.Tuple
)

func BenchmarkSphereIntersect(b *testing.B) {
	s := objects.NewSphere()
	s.SetTransform(core.ScalingMatrix(2, 2, 2))
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkIntersections = s.Intersect(r)
	}
}

func BenchmarkSphereNormalAt(b *testing.B) {
	s := objects.NewSphere()
	s.SetTransform(core.TranslationMatrix(0, 1, 0))
	p := core.NewPoint(0, 1.70711, -0.70711)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkNormal = s.NormalAt(p)
	}
}
