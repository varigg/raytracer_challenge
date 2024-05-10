package shader_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

func TestNewMaterial(t *testing.T) {
	m := shader.NewMaterial()
	assert.Equal(t, m.Color, *core.NewColor(1, 1, 1))
	assert.Equal(t, m.Ambient, 0.1)
	assert.Equal(t, m.Diffuse, 0.9)
	assert.Equal(t, m.Specular, 0.9)
	assert.Equal(t, m.Shininess, 200.0)

}
