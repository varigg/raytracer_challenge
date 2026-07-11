package raytracer

import (
	"math"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/scene"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

var sceneCmd = &cobra.Command{
	Use:     "scene",
	Aliases: []string{"chapter7"},
	Short:   "draws a scene using the camera",
	RunE: func(cmd *cobra.Command, args []string) error {
		wallMaterial := func() *shader.Material {
			return newMaterial(core.NewColor(1, 0.9, 0.9), 0.9, 0)
		}
		wallTransform := func(rotY float64) core.Matrix {
			return core.ScalingMatrix(10, 0.01, 10).RotateX(math.Pi/2).RotateY(rotY).Translate(0, 0, 5)
		}

		floor := newSphere(core.ScalingMatrix(10, 0.01, 10), wallMaterial())
		leftWall := newSphere(wallTransform(-math.Pi/4), wallMaterial())
		rightWall := newSphere(wallTransform(math.Pi/4), wallMaterial())
		middle := newSphere(core.TranslationMatrix(-0.5, 1, 0.5),
			newMaterial(core.NewColor(0.1, 1, 0.5), 0.7, 0.3))
		right := newSphere(core.ScalingMatrix(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5),
			newMaterial(core.NewColor(0.5, 1, 0.1), 0.7, 0.3))
		left := newSphere(core.ScalingMatrix(0.33, 0.33, 0.33).Translate(-1.5, 0.333, -0.75),
			newMaterial(core.NewColor(1, 0.8, 0.1), 0.7, 0.3))

		world := scene.NewWorld()
		world.Add(floor, leftWall, rightWall, middle, right, left)
		world.Light = shader.NewLight(core.NewPoint(-10, 10, -10), core.NewColor(1, 1, 1))

		camera := scene.NewCamera(400, 200, math.Pi/3)
		camera.SetTransform(scene.ViewTransform(core.NewPoint(0, 1.5, -5),
			core.NewPoint(0, 1, 0), core.NewVector(0, 1, 0)))
		return saveCanvas(camera.Render(world), "scene.png")
	},
}

func init() {
	rootCmd.AddCommand(sceneCmd)
}
