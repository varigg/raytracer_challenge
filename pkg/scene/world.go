package scene

import (
	"cmp"
	"slices"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

type World struct {
	Light   *shader.Light
	Objects []objects.Object
}

func NewWorld() *World {
	w := &World{}
	return w
}

func NewDefaultWorld() *World {
	w := &World{}
	w.Objects = make([]objects.Object, 2)
	w.Light = shader.NewLight(core.NewPoint(-10, 10, -10), core.NewColor(1, 1, 1))
	// Outer sphere (book default)
	outerSphere := objects.NewSphere()
	m := shader.NewMaterial()
	m.Color = *core.NewColor(.8, 1, .6)
	m.Diffuse = .7
	m.Specular = .2
	outerSphere.Material = m
	w.Objects[0] = outerSphere
	// Inner sphere (book default: just scaled, default material)
	innerSphere := objects.NewSphere()
	err := innerSphere.SetTransform(core.ScalingMatrix(.5, .5, .5))
	if err != nil {
		return nil
	}
	w.Objects[1] = innerSphere
	return w
}

func (w *World) Intersect(ray *objects.Ray) []objects.Intersection {
	xs := make([]objects.Intersection, 0)
	for _, obj := range w.Objects {
		xs = append(xs, obj.Intersect(ray)...)
	}
	slices.SortFunc(xs, func(a, b objects.Intersection) int {
		return cmp.Compare(a.T, b.T)
	})
	return xs
}

func (w *World) ShadeHit(comps *Computations) *core.Color {
	return w.Light.Lighting(comps.Object.GetMaterial(), comps.Point, comps.EyeV, comps.NormalV)
}

func (w *World) ColorAt(ray *objects.Ray) *core.Color {
	intersections := w.Intersect(ray)
	hit := objects.Hit(intersections)
	if hit == nil {
		return core.NewColor(0, 0, 0)
	}
	comps := PrepareComputations(hit, ray)
	return w.ShadeHit(comps)
}

func (w *World) Add(o ...objects.Object) {
	w.Objects = append(w.Objects, o...)
}
