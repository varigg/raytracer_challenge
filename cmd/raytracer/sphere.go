package raytracer

import (
	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

type pixel struct {
	X int
	Y int
	C *core.Color
}

var sphereShadowCmd = &cobra.Command{
	Use:     "shadow",
	Aliases: []string{"chapter5"},
	Short:   "draws the shadow of a sphere on a wall using raytracing",
	Run: func(cmd *cobra.Command, args []string) {

		// Set up canvas
		pixels := 500
		canvas := core.NewCanvas(pixels, pixels)

		// Ray starting at Z -5
		rayOrigin := core.NewPoint(0, 0, -5)
		// Wall at Z 10
		wallZ := 10.0
		wallSize := 7.0
		pixelSize := wallSize / float64(canvas.Height)
		half := wallSize / 2

		shape := objects.NewSphere()
		shape.SetTransform(core.ScalingMatrix(1, 0.5, 1))

		for y := 0; y < pixels; y += 1 {
			worldY := half - pixelSize*float64(y)
			for x := 0; x < pixels; x += 1 {
				worldX := -half + pixelSize*float64(x)
				// the point on the wall we are shooting the ray at
				position := core.NewPoint(worldX, worldY, wallZ)
				// ray starting at origin with a vector towards the target point
				r := objects.NewRay(rayOrigin, position.Subtract(rayOrigin).Normalize())
				xs := shape.Intersect(r)
				if objects.Hit(xs) != nil {
					canvas.Set(x, y, core.NewColor(0.9, 0, 0))
				}
			}
		}
		canvas.SavePNG("shadow.png")
	},
}

var sphereCmd = &cobra.Command{
	Use:     "sphere",
	Aliases: []string{"chapter6"},
	Short:   "draws a sphere using raytracing",
	Run: func(cmd *cobra.Command, args []string) {

		// Set up canvas
		pixels := 500
		canvas := core.NewCanvas(pixels, pixels)

		// Ray starting at Z -5
		rayOrigin := core.NewPoint(0, 0, -5)
		// Wall at Z 10
		wallZ := 10.0
		wallSize := 7.0
		pixelSize := wallSize / float64(canvas.Height)
		half := wallSize / 2

		shape := objects.NewSphere()
		mat := shader.NewMaterial()
		mat.Color = *core.NewColor(1, 0.2, 1)
		shape.Material = mat
		//shape.SetTransform(core.ScalingMatrix(1, 0.5, 1))
		light := shader.NewLight(core.NewPoint(-10, 10, -10), core.NewColor(1, 1, 1))
		for y := 0; y < pixels; y += 1 {
			worldY := half - pixelSize*float64(y)
			for x := 0; x < pixels; x += 1 {
				worldX := -half + pixelSize*float64(x)
				// the point on the wall we are shooting the ray at
				position := core.NewPoint(worldX, worldY, wallZ)
				// ray starting at origin with a vector towards the target point
				r := objects.NewRay(rayOrigin, position.Subtract(rayOrigin).Normalize())
				xs := shape.Intersect(r)
				hit := objects.Hit(xs)
				if hit != nil {
					p := r.Position(hit.T)
					normal := hit.Object.NormalAt(p)
					eye := r.Direction.Negate()
					color := light.Lighting(mat, p, eye, normal)

					canvas.Set(x, y, color)
				}
			}
		}
		canvas.SavePNG("sphere.png")
	},
}

func init() {
	rootCmd.AddCommand(sphereShadowCmd)
	rootCmd.AddCommand(sphereCmd)
}
