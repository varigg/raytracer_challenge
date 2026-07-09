package raytracer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
)

var rootCmd = &cobra.Command{
	Use:   "raytrace",
	Short: "raytracer - an implementation of the raytracer challenge",
	Run: func(cmd *cobra.Command, args []string) {

	},
	SilenceErrors: true,
	SilenceUsage:  true,
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
	if strings.EqualFold(filepath.Ext(name), ".ppm") {
		return c.SavePPM(name)
	}
	return c.SavePNG(name)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
