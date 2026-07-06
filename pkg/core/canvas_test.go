package core_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
)

func TestNewCanvas(t *testing.T) {
	canvas := core.NewCanvas(10, 20)
	for i := range 10 {
		for j := range 20 {
			assert.True(t, canvas.Get(i, j).Equals(core.NewColor(0, 0, 0)))
		}
	}
}

func TestSetPixel(t *testing.T) {
	canvas := core.NewCanvas(10, 20)
	canvas.Set(2, 3, core.NewColor(1.0, 0, 0))
	assert.True(t, canvas.Get(2, 3).Equals(core.NewColor(1.0, 0, 0)))
}

func TestPPMColorMax(t *testing.T) {
	canvas := core.NewCanvas(5, 3)
	output := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
	canvas.Set(0, 0, core.NewColor(1.5, 0, 0))
	canvas.Set(2, 1, core.NewColor(0, 0.5, 0))
	canvas.Set(4, 2, core.NewColor(-0.5, 0, 1))
	var buf bytes.Buffer
	err := canvas.ToPPM(&buf)
	assert.Nil(t, err)
	assert.Equal(t, output, buf.String())

}

func TestPPMLineLength(t *testing.T) {
	canvas := core.NewCanvas(10, 2)
	output := `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`
	for x := range 10 {
		for y := range 2 {
			canvas.Set(x, y, core.NewColor(1, 0.8, 0.6))
		}
	}
	var buf bytes.Buffer
	err := canvas.ToPPM(&buf)
	assert.Nil(t, err)
	ppm := buf.String()
	assert.Equal(t, output, ppm)
	assert.Equal(t, "\n", string(ppm[len(ppm)-1]))
}
