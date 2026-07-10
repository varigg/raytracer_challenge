package raytracer

import (
	"math"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/scene"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

var simpleSceneCmd = &cobra.Command{
	Use:   "simple-scene",
	Short: "draws a simple scene with two colored spheres",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := objects.NewSphere()
		s.SetTransform(core.TranslationMatrix(-0.5, 0, 0).Times(core.ScalingMatrix(1.5, 1.5, 1.5)))
		mat := shader.NewMaterial()
		mat.Color = core.NewColor(1, 0.2, 1)
		s.Material = mat

		world := scene.NewWorld()
		world.Add(s)
		world.Light = shader.NewLight(core.NewPoint(0, 10, 0), core.NewColor(1, 1, 1))

		camera := scene.NewCamera(500, 500, math.Pi/3)
		camera.Transform = scene.ViewTransform(core.NewPoint(0, 0, -5),
			core.NewPoint(0, 0, 0),
			core.NewVector(0, 1, 0))
		canvas := camera.Render(world)
		return canvas.SavePNG("simple_scene.png")
	},
}

func init() {
	rootCmd.AddCommand(simpleSceneCmd)
}
