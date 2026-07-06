package shader

import "github.com/varigg/raytracer-challenge/pkg/core"

type Material struct {
	Color     core.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() *Material {
	m := &Material{
		Color:     *core.NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
	return m
}
