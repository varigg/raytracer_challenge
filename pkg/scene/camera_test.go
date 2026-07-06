package scene_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/scene"
)

func TestTransform(t *testing.T) {
	// Default orientation
	m := scene.ViewTransform(core.NewPoint(0, 0, 0), core.NewPoint(0, 0, -1), core.NewVector(0, 1, 0))
	assert.Equal(t, core.Identity4(), m)
	// backwards
	m = scene.ViewTransform(core.NewPoint(0, 0, 0), core.NewPoint(0, 0, 1), core.NewVector(0, 1, 0))
	assert.Equal(t, core.ScalingMatrix(-1, 1, -1), m)
	// move away
	m = scene.ViewTransform(core.NewPoint(0, 0, 8), core.NewPoint(0, 0, 0), core.NewVector(0, 1, 0))
	assert.Equal(t, core.TranslationMatrix(0, 0, -8), m)
	// arbitrary position
	m = scene.ViewTransform(core.NewPoint(1, 3, 2), core.NewPoint(4, -2, 8), core.NewVector(1, 1, 0))
	transform := core.NewMatrix([][]float64{
		{-0.50709, 0.50709, 0.67612, -2.36643},
		{0.76772, 0.60609, 0.12122, -2.82843},
		{-0.35857, 0.59761, -0.71714, 0.00000},
		{0.00000, 0.00000, 0.00000, 1.00000},
	})
	//assert.Equal(t, transform, m)
	assert.True(t, transform.Equals(m))
}

func TestNewCamera(t *testing.T) {
	c := scene.NewCamera(160, 120, math.Pi/2)
	assert.Equal(t, 160, c.HSize)
	assert.Equal(t, 120, c.VSize)
	assert.Equal(t, math.Pi/2, c.FOV)
	assert.Equal(t, core.Identity4(), c.Transform)
}
func TestPixelSize(t *testing.T) {
	c := scene.NewCamera(200, 125, math.Pi/2)
	assert.Equal(t, 0.01, c.PixelSize)
	c = scene.NewCamera(125, 200, math.Pi/2)
	assert.Equal(t, 0.01, c.PixelSize)
}

func TestRayForPixel(t *testing.T) {
	// ray through center
	c := scene.NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)
	assert.Equal(t, core.NewPoint(0, 0, 0), r.Origin)
	assert.Equal(t, core.NewVector(0, 0, -1), r.Direction)
	// ray through corner
	r = c.RayForPixel(0, 0)
	assert.Equal(t, core.NewPoint(0, 0, 0), r.Origin)
	assert.True(t, core.NewVector(0.66519, 0.33259, -0.66851).Equals(r.Direction))
	// ray with transformed camera
	c.Transform = core.RotationMatrixY(math.Pi / 4).Times(core.TranslationMatrix(0, -2, 5))
	r = c.RayForPixel(100, 50)
	assert.Equal(t, core.NewPoint(0, 2, -5), r.Origin)
	assert.True(t, core.NewVector(math.Sqrt2/2, 0, -math.Sqrt2/2).Equals(r.Direction))
}

func TestRender(t *testing.T) {
	w := scene.NewDefaultWorld()
	c := scene.NewCamera(11, 11, math.Pi/2)
	from := core.NewPoint(0, 0, -5)
	to := core.NewPoint(0, 0, 0)
	up := core.NewVector(0, 1, 0)
	c.Transform = scene.ViewTransform(from, to, up)
	image := c.Render(w)
	assert.True(t, core.NewColor(0.38066, 0.47583, 0.2855).Equals(image.Get(5, 5)))
}
