package raytracer

import (
	"math"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
)

var clockCmd = &cobra.Command{
	Use:     "clock",
	Aliases: []string{"chapter4"},
	Short:   "draws a clock face using matrix operations",
	Run: func(cmd *cobra.Command, args []string) {
		r := core.RotationMatrixY(2.0 * math.Pi / 12.0)
		var x, y int
		current := core.NewPoint(0.0, 0.0, 1.0)

		c := core.NewCanvas(500, 500)
		color := core.NewColor(1.0, 0.0, 0.0)
		radius := 3.0 / 8.0 * 500.0

		for _ = range 12 {
			x = int(current.X*radius + 250)
			y = int(current.Z*radius + 250)
			DrawSquare(c, x, y, color)

			current = r.MultiplyWithTuple(current)
		}
		c.SavePPM("clock.ppm")
		c.SavePNG("clock.png")
	},
}

func init() {
	rootCmd.AddCommand(clockCmd)
}
