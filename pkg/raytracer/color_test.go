package raytracer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/raytracer"
)

func TestNewColor(t *testing.T) {
	c := raytracer.NewColor(-0.5, 0.4, 1.7)
	assert.Equal(t, -0.5, c.Red())
	assert.Equal(t, 0.4, c.Green())
	assert.Equal(t, 1.7, c.Blue())
}

func TestAddColor(t *testing.T) {
	c1 := raytracer.NewColor(0.9, 0.6, 0.75)
	c2 := raytracer.NewColor(0.7, 0.1, 0.25)
	assert.Equal(t, raytracer.NewColor(1.6, 0.7, 1.0), c1.Add(c2))
}

func TestSubtractColor(t *testing.T) {
	c1 := raytracer.NewColor(0.9, 0.6, 0.75)
	c2 := raytracer.NewColor(0.7, 0.1, 0.25)
	assert.True(t, raytracer.NewColor(0.2, 0.5, 0.5).Equals(c1.Subtract(c2)))
}

func TestMultiplyColor(t *testing.T) {
	c := raytracer.NewColor(0.2, 0.3, 0.4)
	assert.True(t, raytracer.NewColor(0.4, 0.6, 0.8).Equals(c.Multiply(2)))
}

func TestHadamardProduct(t *testing.T) {
	c1 := raytracer.NewColor(1, 0.2, 0.4)
	c2 := raytracer.NewColor(0.9, 1, 0.1)
	assert.True(t, raytracer.NewColor(0.9, 0.2, 0.04).Equals(c1.HadamardProduct(c2)))
}

func TestFloatToColorValue(t *testing.T) {
	assert.Equal(t, 255, raytracer.ConvertFloatToColorValue(1.0, 255))
	assert.Equal(t, 0, raytracer.ConvertFloatToColorValue(0.0, 255))
	assert.Equal(t, 0, raytracer.ConvertFloatToColorValue(-1.0, 255))
	assert.Equal(t, 255, raytracer.ConvertFloatToColorValue(1.5, 255))
	assert.Equal(t, 128, raytracer.ConvertFloatToColorValue(.5, 255))
}

func TestToRGB(t *testing.T) {
	c1 := raytracer.NewColor(1, 0.2, 0.4)
	r, g, b := c1.ToRGB(255)
	assert.Equal(t, 255, r)
	assert.Equal(t, 51, g)
	assert.Equal(t, 102, b)
}
