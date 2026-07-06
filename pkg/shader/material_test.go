package shader_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

// Scenario: The default material
//
//	Given m ← material()
//	Then m.color = color(1, 1, 1)
//	  And m.ambient = 0.1
//	  And m.diffuse = 0.9
//	  And m.specular = 0.9
//	  And m.shininess = 200.0
func TestNewMaterial(t *testing.T) {
	m := shader.NewMaterial()
	assert.Equal(t, m.Color, *core.NewColor(1, 1, 1))
	assert.Equal(t, m.Ambient, 0.1)
	assert.Equal(t, m.Diffuse, 0.9)
	assert.Equal(t, m.Specular, 0.9)
	assert.Equal(t, m.Shininess, 200.0)

}
