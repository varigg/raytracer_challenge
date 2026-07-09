package raytracer

import (
	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

var sphereShadowCmd = &cobra.Command{
	Use:     "shadow",
	Aliases: []string{"chapter5"},
	Short:   "draws the shadow of a sphere on a wall using raytracing",
	RunE: func(cmd *cobra.Command, args []string) error {
		shape := objects.NewSphere()
		shape.SetTransform(core.ScalingMatrix(1, 0.5, 1))
		canvas := renderOnWall(shape, func(_ *objects.Intersection, _ *objects.Ray) core.Color {
			return core.NewColor(0.9, 0, 0)
		})
		return saveCanvas(canvas, "shadow.png")
	},
}

var sphereCmd = &cobra.Command{
	Use:     "sphere",
	Aliases: []string{"chapter6"},
	Short:   "draws a sphere using raytracing",
	RunE: func(cmd *cobra.Command, args []string) error {
		shape := objects.NewSphere()
		shape.SetMaterial(newMaterial(core.NewColor(1, 0.2, 1), 0.9, 0.9))
		light := shader.NewLight(core.NewPoint(-10, 10, -10), core.NewColor(1, 1, 1))
		canvas := renderOnWall(shape, func(hit *objects.Intersection, ray *objects.Ray) core.Color {
			point := ray.Position(hit.T)
			normal := hit.Object.NormalAt(point)
			eye := ray.Direction.Negate()
			return light.Lighting(hit.Object.Material(), point, eye, normal)
		})
		return saveCanvas(canvas, "sphere.png")
	},
}

func init() {
	rootCmd.AddCommand(sphereShadowCmd)
	rootCmd.AddCommand(sphereCmd)
}
