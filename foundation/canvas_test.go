package foundation_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytacer-challenge/foundation"
)

func TestCanvas(t *testing.T) {
	c := foundation.NewCanvas(10, 20)
	assert.Equal(t, 10, c.Width)
	assert.Equal(t, 20, c.Height)
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			assert.Equal(t, foundation.NewColor(0, 0, 0), c.GetPixel(i, j))
		}
	}
}

func TestPixels(t *testing.T) {
	c := foundation.NewCanvas(10, 20)
	red := foundation.NewColor(1, 0, 0)
	c.SetPixel(2, 3, red)
	assert.Equal(t, red, c.GetPixel(2, 3))
	assert.Equal(t, foundation.NewColor(0, 0, 0), c.GetPixel(3, 2))
}

func TestPPMHeader(t *testing.T) {
	c := foundation.NewCanvas(5, 3)
	ppm := c.ToPPM()
	lines := strings.Split(ppm, "\n")
	assert.Equal(t, "P3", lines[0])
	assert.Equal(t, "5 3", lines[1])
	assert.Equal(t, "255", lines[2])
}
func TestPPM(t *testing.T) {
	canvas := foundation.NewCanvas(5, 3)
	c1 := foundation.NewColor(1.5, 0, 0)
	c2 := foundation.NewColor(0, 0.5, 0)
	c3 := foundation.NewColor(-0.5, 0, 1)
	canvas.SetPixel(0, 0, c1)
	canvas.SetPixel(2, 1, c2)
	canvas.SetPixel(4, 2, c3)
	ppm := canvas.ToPPM()
	lines := strings.Split(ppm, "\n")
	t.Log(ppm)
	assert.Equal(t, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", lines[3])
	assert.Equal(t, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0", lines[4])
	assert.Equal(t, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", lines[5])
}

func TestPPMLongLines(t *testing.T) {
	canvas := foundation.NewCanvas(10, 2)
	color := foundation.NewColor(1, 0.8, 0.6)
	for i := 0; i < 10; i++ {
		for j := 0; j < 2; j++ {
			canvas.SetPixel(i, j, color)
		}
	}
	ppm := canvas.ToPPM()
	lines := strings.Split(ppm, "\n")
	t.Log(ppm)
	assert.Equal(t, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", lines[3])
	assert.Equal(t, "153 255 204 153 255 204 153 255 204 153 255 204 153", lines[4])
	assert.Equal(t, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", lines[5])
	assert.Equal(t, "153 255 204 153 255 204 153 255 204 153 255 204 153", lines[6])
	assert.Equal(t, "", lines[len(lines)-1])
}
