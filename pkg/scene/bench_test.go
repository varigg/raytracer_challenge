package scene_test

import (
	"math"
	"testing"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/scene"
)

var sinkCanvas *core.Canvas

func BenchmarkCameraRender(b *testing.B) {
	w := scene.NewDefaultWorld()
	c := scene.NewCamera(100, 50, math.Pi/3)
	c.SetTransform(scene.ViewTransform(
		core.NewPoint(0, 0, -5), core.NewPoint(0, 0, 0), core.NewVector(0, 1, 0)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkCanvas = c.Render(w)
	}
}
