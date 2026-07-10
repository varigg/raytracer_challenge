package shader_test

import (
	"testing"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

var sinkColor *core.Color

func BenchmarkLighting(b *testing.B) {
	m := shader.NewMaterial()
	light := shader.NewLight(core.NewPoint(0, 0, -10), core.NewColor(1, 1, 1))
	position := core.NewPoint(0, 0, 0)
	eye := core.NewVector(0, 0, -1)
	normal := core.NewVector(0, 0, -1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkColor = light.Lighting(m, position, eye, normal)
	}
}
