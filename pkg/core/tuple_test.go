package core_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
)

func TestNewTuple(t *testing.T) {
	tuple := core.NewTuple(4.3, -4.2, 3.1, 1.0)
	assert.NotNil(t, tuple)
	assert.True(t, tuple.IsPoint())
	assert.False(t, tuple.IsVector())
}

func TestNewVector(t *testing.T) {
	v := core.NewVector(4, -4, 3)
	assert.True(t, v.IsVector())
}

func TestNewPoint(t *testing.T) {
	v := core.NewPoint(4, -4, 3)
	assert.True(t, v.IsPoint())
}

func TestAdd(t *testing.T) {
	t1 := core.NewTuple(3, -2, 5, 1)
	t2 := core.NewTuple(-2, 3, 1, 0)
	assert.Equal(t, core.NewTuple(1, 1, 6, 1), t1.Add(t2))
}

func TestSubtractPointPoint(t *testing.T) {
	t1 := core.NewPoint(3, 2, 1)
	t2 := core.NewPoint(5, 6, 7)
	assert.Equal(t, core.NewVector(-2, -4, -6), t1.Subtract(t2))
}

func TestSubtractPointVector(t *testing.T) {
	t1 := core.NewPoint(3, 2, 1)
	t2 := core.NewVector(5, 6, 7)
	assert.Equal(t, core.NewPoint(-2, -4, -6), t1.Subtract(t2))
}

func TestSubtractVectorVector(t *testing.T) {
	t1 := core.NewVector(3, 2, 1)
	t2 := core.NewVector(5, 6, 7)
	assert.Equal(t, core.NewVector(-2, -4, -6), t1.Subtract(t2))
}

func TestSubtractVectorZeroVector(t *testing.T) {
	zero := core.NewVector(0, 0, 0)
	t2 := core.NewVector(1, -2, 3)
	assert.Equal(t, core.NewVector(-1, 2, -3), zero.Subtract(t2))
}

func TestNegate(t *testing.T) {
	tuple := core.NewTuple(1, -2, 3, -4)
	assert.Equal(t, core.NewTuple(-1, 2, -3, 4), tuple.Negate())
}

func TestDivide(t *testing.T) {
	tuple := core.NewTuple(1, -2, 3, -4)
	assert.Equal(t, core.NewTuple(0.5, -1, 1.5, -2), tuple.Divide(2))

}

func TestMultiply(t *testing.T) {
	tuple := core.NewTuple(1, -2, 3, -4)
	assert.Equal(t, core.NewTuple(3.5, -7, 10.5, -14), tuple.Multiply(3.5))
	assert.Equal(t, core.NewTuple(0.5, -1, 1.5, -2), tuple.Multiply(0.5))
}

func TestMagnitue(t *testing.T) {
	tuple := core.NewVector(1, 0, 0)
	assert.Equal(t, 1.0, tuple.Magnitude())
	tuple = core.NewVector(0, 1, 0)
	assert.Equal(t, 1.0, tuple.Magnitude())
	tuple = core.NewVector(0, 0, 1)
	assert.Equal(t, 1.0, tuple.Magnitude())
	tuple = core.NewVector(1, 2, 3)
	assert.Equal(t, math.Sqrt(14), tuple.Magnitude())
	tuple = core.NewVector(-1, -2, -3)
	assert.Equal(t, math.Sqrt(14), tuple.Magnitude())
}

func TestNormalize(t *testing.T) {
	vector := core.NewVector(4, 0, 0)
	assert.Equal(t, core.NewVector(1, 0, 0), vector.Normalize())
	vector = core.NewVector(1, 2, 3)
	assert.Equal(t, core.NewVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)), vector.Normalize())
	assert.Equal(t, 1.0, vector.Normalize().Magnitude())
}

func TestDot(t *testing.T) {
	v1 := core.NewVector(1, 2, 3)
	v2 := core.NewVector(2, 3, 4)
	assert.Equal(t, 20.0, v1.Dot(v2))
}

func TestCross(t *testing.T) {
	v1 := core.NewVector(1, 2, 3)
	v2 := core.NewVector(2, 3, 4)
	assert.Equal(t, core.NewVector(-1, 2, -1), v1.Cross(v2))
	assert.Equal(t, core.NewVector(1, -2, 1), v2.Cross(v1))
}
