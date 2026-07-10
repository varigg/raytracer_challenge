package objects

import (
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

type Object interface {
	Intersect(*Ray) []Intersection
	NormalAt(core.Tuple) core.Tuple
	Material() *shader.Material
}
