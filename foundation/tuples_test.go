package foundation_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytacer-challenge/foundation"
)

func TestPoint(t *testing.T) {
	tuple := foundation.Tuple{4.3, -4.2, 3.1, 1}
	assert.True(t, tuple.IsPoint())
	assert.False(t, tuple.IsVector())
	assert.True(t, foundation.NewPoint(4.3, -4.2, 3.1).IsPoint())

}

func TestVector(t *testing.T) {
	tuple := foundation.NewTuple(4.3, -4.2, 3.1, 0)
	assert.True(t, tuple.IsVector())
	assert.False(t, tuple.IsPoint())
	assert.True(t, foundation.NewVector(4.3, -4.2, 3.1).IsVector())
}

func TestAdd(t *testing.T) {
	a1 := foundation.NewTuple(3, -2, 5, 1)
	a2 := foundation.NewTuple(-2, 3, 1, 0)
	assert.True(t, a1.Plus(a2).Equals(foundation.NewTuple(1, 1, 6, 1)))
}

func TestSubstractPoints(t *testing.T) {
	p1 := foundation.NewPoint(3, 2, 1)
	p2 := foundation.NewPoint(5, 6, 7)
	assert.True(t, p1.Minus(p2).Equals(foundation.NewVector(-2, -4, -6)))
}

func TestSubstractVectorFromPoint(t *testing.T) {
	p1 := foundation.NewPoint(3, 2, 1)
	p2 := foundation.NewVector(5, 6, 7)
	assert.True(t, p1.Minus(p2).Equals(foundation.NewPoint(-2, -4, -6)))
}

func TestSubstractVectors(t *testing.T) {
	v1 := foundation.NewVector(3, 2, 1)
	v2 := foundation.NewVector(5, 6, 7)
	assert.True(t, v1.Minus(v2).Equals(foundation.NewVector(-2, -4, -6)))
}

func TestSubstractFromZeroVector(t *testing.T) {
	zero := foundation.NewVector(0, 0, 0)
	v := foundation.NewVector(1, -2, 3)
	assert.True(t, zero.Minus(v).Equals(foundation.NewVector(-1, 2, -3)))
}

func TestNegateVector(t *testing.T) {
	v := foundation.NewVector(1, -2, 3)
	assert.True(t, v.Negate().Equals(foundation.NewVector(-1, 2, -3)))
}

func TestScalarMultiplication(t *testing.T) {
	tuple := foundation.NewTuple(1, -2, 3, -4)
	assert.Equal(t, foundation.NewTuple(3.5, -7, 10.5, -14), tuple.Times(3.5))
	assert.Equal(t, foundation.NewTuple(.5, -1, 1.5, -2), tuple.Times(.5))
}

func TestScalarDivision(t *testing.T) {
	tuple := foundation.NewTuple(1, -2, 3, -4)
	assert.Equal(t, foundation.NewTuple(.5, -1, 1.5, -2), tuple.DivideBy(2))
}
func TestMagnitude(t *testing.T) {
	assert.Equal(t, float64(1), foundation.NewVector(1, 0, 0).Magnitude())
	assert.Equal(t, float64(1), foundation.NewVector(0, 1, 0).Magnitude())
	assert.Equal(t, float64(1), foundation.NewVector(0, 0, 1).Magnitude())
	assert.Equal(t, float64(math.Sqrt(14)), foundation.NewVector(1, 2, 3).Magnitude())
	assert.Equal(t, float64(math.Sqrt(14)), foundation.NewVector(-1, -2, -3).Magnitude())
}

func TestNormalize(t *testing.T) {
	assert.Equal(t, foundation.NewVector(1, 0, 0), foundation.NewVector(4, 0, 0).Normalize())
	assert.True(t, foundation.NewVector(0.26726, .53452, .80178).Equals(foundation.NewVector(1, 2, 3).Normalize()))
	assert.Equal(t, float64(1), foundation.NewVector(1, 2, 3).Normalize().Magnitude())
}

func TestDotProduct(t *testing.T) {
	assert.Equal(t, float64(20), foundation.NewVector(1, 2, 3).DotProduct(foundation.NewVector(2, 3, 4)))
}

func TestCrossProduct(t *testing.T) {
	a := foundation.NewVector(1, 2, 3)
	b := foundation.NewVector(2, 3, 4)
	cp, err := a.CrossProduct(b)
	assert.Nil(t, err)
	assert.Equal(t, foundation.NewVector(-1, 2, -1), cp)
	cp, err = b.CrossProduct(a)
	assert.Nil(t, err)
	assert.Equal(t, foundation.NewVector(1, -2, 1), cp)
	p := foundation.NewPoint(1, 1, 1)
	cp, err = a.CrossProduct(p)
	assert.NotNil(t, err)
	assert.Nil(t, cp)
}
