package objects

import (
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

type Object interface {
	Intersect(*Ray) []Intersection
	NormalAt(core.Tuple) core.Tuple
	GetMaterial() *shader.Material
	SetTransform(*core.Matrix) error
	GetTransform() *core.Matrix
	GetInverseTransform() *core.Matrix
}
