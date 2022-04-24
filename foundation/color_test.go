package foundation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytacer-challenge/foundation"
)

func TestColorsAdd(t *testing.T) {
	c1 := foundation.NewColor(0.9, 0.6, 0.75)
	c2 := foundation.NewColor(0.7, .1, .25)
	assert.Equal(t, foundation.NewColor(1.6, 0.7, 1.0), c1.Plus(c2))
}
func TestColorsSubstract(t *testing.T) {
	c1 := foundation.NewColor(0.9, 0.6, 0.75)
	c2 := foundation.NewColor(0.7, .1, .25)
	assert.True(t, foundation.NewColor(0.2, 0.5, 0.5).Equals(c1.Minus(c2)))
}
func TestColorsMultiply(t *testing.T) {
	c1 := foundation.NewColor(0.2, 0.3, 0.4)
	assert.True(t, foundation.NewColor(0.4, 0.6, 0.8).Equals(c1.Times(2)))
}

func TestColorsBlend(t *testing.T) {
	c1 := foundation.NewColor(1, 0.2, 0.4)
	c2 := foundation.NewColor(0.9, 1, .1)
	assert.True(t, foundation.NewColor(0.9, 0.2, 0.04).Equals(c1.Blend(c2)))
}
