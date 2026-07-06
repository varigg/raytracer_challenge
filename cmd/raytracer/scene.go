package raytracer

import (
	"math"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/scene"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

var sceneCmd = &cobra.Command{
	Use:     "scene",
	Aliases: []string{"chapter7"},
	Short:   "draws a scene using the camera",
	Run: func(cmd *cobra.Command, args []string) {
		floor := objects.NewSphere()
		floor.SetTransform(core.ScalingMatrix(10, 0.01, 10))
		floor.Material = shader.NewMaterial()
		floor.Material.Color = *core.NewColor(1, 0.9, 0.9)
		floor.Material.Specular = 0
		floor.Material.Ambient = 0.1
		floor.Material.Diffuse = 0.9

		leftWall := objects.NewSphere()
		leftWall.SetTransform(
			core.TranslationMatrix(0, 0, 5).Times(
				core.RotationMatrixY(-math.Pi / 4).Times(
					core.RotationMatrixX(math.Pi / 2).Times(
						core.ScalingMatrix(10, .01, 10)))))
		leftWall.Material = floor.Material
		leftWall.Material.Ambient = 0.1
		leftWall.Material.Diffuse = 0.9

		rightWall := objects.NewSphere()
		rightWall.SetTransform(
			core.TranslationMatrix(0, 0, 5).Times(
				core.RotationMatrixY(math.Pi / 4).Times(
					core.RotationMatrixX(math.Pi / 2).Times(
						core.ScalingMatrix(10, .01, 10)))))
		rightWall.Material = floor.Material
		rightWall.Material.Ambient = 0.1
		rightWall.Material.Diffuse = 0.9

		middle := objects.NewSphere()
		middle.SetTransform(core.TranslationMatrix(-0.5, 1, 0.5))
		middle.Material = shader.NewMaterial()
		middle.Material.Color = *core.NewColor(0.1, 1, 0.5)
		middle.Material.Diffuse = 0.7
		middle.Material.Specular = 0.3
		middle.Material.Ambient = 0.1

		right := objects.NewSphere()
		right.SetTransform(core.TranslationMatrix(1.5, 0.5, -0.5).Times(core.ScalingMatrix(0.5, 0.5, 0.5)))
		right.Material = shader.NewMaterial()
		right.Material.Color = *core.NewColor(0.5, 1, 0.1)
		right.Material.Diffuse = 0.7
		right.Material.Specular = 0.3
		right.Material.Ambient = 0.1

		left := objects.NewSphere()
		left.Material = shader.NewMaterial()
		left.Material.Color = *core.NewColor(1, 0.8, 0.1)
		left.SetTransform(core.TranslationMatrix(-1.5, 0.333, -0.75).Times(core.ScalingMatrix(0.33, 0.33, 0.33)))
		left.Material.Color = *core.NewColor(1, 0.8, 0.1)
		left.Material.Diffuse = 0.7
		left.Material.Specular = 0.3
		left.Material.Ambient = 0.1

		world := scene.NewWorld()
		world.Add(floor, leftWall, rightWall, middle, right, left)

		world.Light = shader.NewLight(core.NewPoint(-10, 10, -10), core.NewColor(1, 1, 1))
		camera := scene.NewCamera(400, 200, math.Pi/3)
		camera.Transform = scene.ViewTransform(core.NewPoint(0, 1.5, -5),
			core.NewPoint(0, 1, 0),
			core.NewVector(0, 1, 0))
		canvas := camera.Render(world)
		canvas.SavePNG("scene.png")

	},
}

func init() {
	rootCmd.AddCommand(sceneCmd)
}
