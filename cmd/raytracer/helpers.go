package raytracer

import (
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

func newMaterial(c core.Color, diffuse, specular float64) *shader.Material {
	m := shader.NewMaterial()
	m.Color = c
	m.Diffuse = diffuse
	m.Specular = specular
	return m
}

func newSphere(transform core.Matrix, mat *shader.Material) *objects.Sphere {
	s := objects.NewSphere()
	s.SetTransform(transform)
	s.SetMaterial(mat)
	return s
}

// renderOnWall shoots a ray from a fixed origin at every pixel of a 7x7 wall
// at z=10 and paints pixels where shape is hit, using shade for the color.
func renderOnWall(shape *objects.Sphere, shade func(hit *objects.Intersection, ray *objects.Ray) core.Color) *core.Canvas {
	const pixels = 500
	const wallZ, wallSize = 10.0, 7.0
	canvas := core.NewCanvas(pixels, pixels)
	rayOrigin := core.NewPoint(0, 0, -5)
	pixelSize := wallSize / float64(pixels)
	half := wallSize / 2
	for y := 0; y < pixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < pixels; x++ {
			worldX := -half + pixelSize*float64(x)
			target := core.NewPoint(worldX, worldY, wallZ)
			ray := objects.NewRay(rayOrigin, target.Subtract(rayOrigin).Normalize())
			if hit := objects.Hit(shape.Intersect(ray)); hit != nil {
				canvas.Set(x, y, shade(hit, ray))
			}
		}
	}
	return canvas
}
