package raytracer

import (
	"math"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/scene"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

var simpleSceneCmd = &cobra.Command{
	Use:   "simple-scene",
	Short: "draws a simple scene with two colored spheres",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := newSphere(core.ScalingMatrix(1.5, 1.5, 1.5).Translate(-0.5, 0, 0),
			newMaterial(core.NewColor(1, 0.2, 1), 0.9, 0.9))

		world := scene.NewWorld()
		world.Add(s)
		world.Light = shader.NewLight(core.NewPoint(0, 10, 0), core.NewColor(1, 1, 1))

		camera := scene.NewCamera(500, 500, math.Pi/3)
		camera.SetTransform(scene.ViewTransform(core.NewPoint(0, 0, -5),
			core.NewPoint(0, 0, 0),
			core.NewVector(0, 1, 0)))
		return saveCanvas(camera.Render(world), "simple_scene.png")
	},
}

func init() {
	rootCmd.AddCommand(simpleSceneCmd)
}
