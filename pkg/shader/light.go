package shader

import (
	"math"

	"github.com/varigg/raytracer-challenge/pkg/core"
)

type Light struct {
	Position  *core.Tuple
	Intensity *core.Color
}

func NewLight(position *core.Tuple, color *core.Color) *Light {
	l := &Light{
		Position:  position,
		Intensity: color,
	}
	return l
}

func (l *Light) Lighting(m *Material, point, eyeV, normalV *core.Tuple) *core.Color {
	effectiveColor := m.Color.HadamardProduct(l.Intensity)
	lightV := l.Position.Subtract(point).Normalize()
	ambient := effectiveColor.Multiply(m.Ambient)
	lightDotNormal := lightV.Dot(normalV)
	var diffuse, specular *core.Color
	black := core.NewColor(0, 0, 0)
	if lightDotNormal < 0 {
		diffuse = black
		specular = black
	} else {
		diffuse = effectiveColor.Multiply(m.Diffuse).Multiply(lightDotNormal)
		reflectV := lightV.Reflect(normalV)
		reflectDotEye := reflectV.Negate().Dot(eyeV)
		if reflectDotEye <= 0 {
			specular = black
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.Multiply(m.Specular).Multiply(factor)
		}
	}
	return ambient.Add(diffuse).Add(specular)
}
