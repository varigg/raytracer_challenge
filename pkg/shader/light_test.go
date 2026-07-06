package shader_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

func TestNew(t *testing.T) {
	l := shader.NewLight(core.NewPoint(0, 0, 0), core.NewColor(1, 1, 1))
	assert.Equal(t, l.Position, core.NewPoint(0, 0, 0))
	assert.Equal(t, l.Intensity, core.NewColor(1, 1, 1))
}

func TestLighting(t *testing.T) {
	m := shader.NewMaterial()
	pos := core.NewPoint(0, 0, 0)
	// Lighting with the eye between the light and the surface
	light := shader.NewLight(core.NewPoint(0, 0, -10), core.NewColor(1, 1, 1))
	eyeV := core.NewVector(0, 0, -1)
	normalV := core.NewVector(0, 0, -1)
	result := light.Lighting(m, pos, eyeV, normalV)
	assert.Equal(t, core.NewColor(1.9, 1.9, 1.9), result)
	assert.True(t, result.Equals(core.NewColor(1.9, 1.9, 1.9)))
	// Lighting with the eye between light and surface, eye offset 45 degree
	eyeV = core.NewVector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV = core.NewVector(0, 0, -1)
	result = light.Lighting(m, pos, eyeV, normalV)
	assert.True(t, result.Equals(core.NewColor(1.0, 1.0, 1.0)))
	// Lighting with eye opposite surface, light offset 45 degrees
	light = shader.NewLight(core.NewPoint(0, 10, -10), core.NewColor(1, 1, 1))
	eyeV = core.NewVector(0, 0, -1)
	normalV = core.NewVector(0, 0, -1)
	result = light.Lighting(m, pos, eyeV, normalV)
	assert.True(t, result.Equals(core.NewColor(0.7364, 0.7364, 0.7364)))
	// Lighting with eye in the path of the reflection vector
	light = shader.NewLight(core.NewPoint(0, 10, -10), core.NewColor(1, 1, 1))
	eyeV = core.NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV = core.NewVector(0, 0, -1)
	result = light.Lighting(m, pos, eyeV, normalV)
	assert.True(t, result.Equals(core.NewColor(1.6364, 1.6364, 1.6364)))
	// Lighting with the light behind the surface
	light = shader.NewLight(core.NewPoint(0, 0, 10), core.NewColor(1, 1, 1))
	eyeV = core.NewVector(0, 0, -1)
	normalV = core.NewVector(0, 0, -1)
	result = light.Lighting(m, pos, eyeV, normalV)
	assert.True(t, result.Equals(core.NewColor(0.1, 0.1, 0.1)))

}
