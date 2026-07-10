package raytracer

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
)

var rootCmd = &cobra.Command{
	Use:   "raytrace",
	Short: "raytracer - an implementation of the raytracer challenge",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var outputFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "",
		"output image file; .ppm extension selects PPM, anything else PNG (default per command)")
}

// saveCanvas writes c to --output if set, else to defaultName, picking the
// format from the file extension.
func saveCanvas(c *core.Canvas, defaultName string) error {
	name := outputFile
	if name == "" {
		name = defaultName
	}
	if strings.HasSuffix(name, ".ppm") {
		return c.SavePPM(name)
	}
	return c.SavePNG(name)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
