package core_test

import (
	"testing"

	"github.com/varigg/raytracer-challenge/pkg/core"
)

var benchMatrix = core.NewMatrix([][]float64{
	{-5, 2, 6, -8},
	{1, -5, 1, 8},
	{7, 7, -6, -7},
	{1, -3, 7, 4},
})

var (
	sinkMatrix core.Matrix
	sinkTuple  core.Tuple
)

func BenchmarkMatrixInvert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkMatrix = benchMatrix.Invert()
	}
}

func BenchmarkMatrixTimes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkMatrix = benchMatrix.Times(benchMatrix)
	}
}

func BenchmarkTupleOps(b *testing.B) {
	v := core.NewVector(1, -1, 2)
	n := core.NewVector(0, 1, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkTuple = v.Add(v).Multiply(0.5).Reflect(n).Normalize()
	}
}
