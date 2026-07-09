package raytracer

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/core"
)

var wind string
var gravity string
var velocity string
var origin string

var projectileCmd = &cobra.Command{
	Use:     "projectile",
	Aliases: []string{"chapter1"},
	Short:   "computes how far a projectile flies until it hits y=0",
	RunE: func(cmd *cobra.Command, args []string) error {
		position, err := core.NewPointFromString(origin)
		if err != nil {
			return fmt.Errorf("--origin: %w", err)
		}
		vel, err := core.NewVectorFromString(velocity)
		if err != nil {
			return fmt.Errorf("--velocity: %w", err)
		}
		grav, err := core.NewVectorFromString(gravity)
		if err != nil {
			return fmt.Errorf("--gravity: %w", err)
		}
		wnd, err := core.NewVectorFromString(wind)
		if err != nil {
			return fmt.Errorf("--wind: %w", err)
		}
		p := projectile{position: position, velocity: vel}
		env := environment{gravity: grav, wind: wnd}
		distance := 0
		for p.position.Y > 0 {
			fmt.Println(p.position.X)
			p = tick(p, env)
			distance += 1
		}
		fmt.Println(distance)
		return nil
	},
}

var projectileGraphCmd = &cobra.Command{
	Use:     "projectile-graph",
	Aliases: []string{"chapter2"},
	Short:   "plots trajectory of a projectile",
	RunE: func(cmd *cobra.Command, args []string) error {
		position, err := core.NewPointFromString(origin)
		if err != nil {
			return fmt.Errorf("--origin: %w", err)
		}
		vel, err := core.NewVectorFromString(velocity)
		if err != nil {
			return fmt.Errorf("--velocity: %w", err)
		}
		grav, err := core.NewVectorFromString(gravity)
		if err != nil {
			return fmt.Errorf("--gravity: %w", err)
		}
		wnd, err := core.NewVectorFromString(wind)
		if err != nil {
			return fmt.Errorf("--wind: %w", err)
		}
		p := projectile{position: position, velocity: vel}
		env := environment{gravity: grav, wind: wnd}
		positions := make([]core.Tuple, 0)
		maxX, maxY := 0.0, 0.0
		positions = append(positions, p.position)
		for p.position.Y > 0 {
			if p.position.X > maxX {
				maxX = p.position.X
			}
			if p.position.Y > maxY {
				maxY = p.position.Y
			}
			positions = append(positions, p.position)
			p = tick(p, env)
		}
		canvas := core.NewCanvas(int(maxX)+10, int(maxY)+10)
		fmt.Printf("%dx%d", canvas.Width, canvas.Height)
		red := core.NewColor(1.0, 0, 0)
		for i := range positions {
			x := int(positions[i].X) + 5
			y := canvas.Height - int(positions[i].Y) - 5
			canvas.DrawSquare(x, y, red)
		}
		return saveCanvas(canvas, "canvas.ppm")
	},
}

type projectile struct {
	position core.Tuple
	velocity core.Tuple
}

type environment struct {
	gravity core.Tuple
	wind    core.Tuple
}

func init() {
	projectileCmd.Flags().StringVarP(&origin, "origin", "o", "0,1,0", "the starting point")
	projectileCmd.Flags().StringVarP(&velocity, "velocity", "v", "1,1,0", "the initial velocity")
	projectileCmd.Flags().StringVarP(&gravity, "gravity", "g", "0,-0.1,0", "the gravity vector")
	projectileCmd.Flags().StringVarP(&wind, "wind", "w", "-0.01,0,0", "the wind vector")
	projectileGraphCmd.Flags().StringVarP(&origin, "origin", "o", "0,1,0", "the starting point")
	projectileGraphCmd.Flags().StringVarP(&velocity, "velocity", "v", "1,1,0", "the initial velocity")
	projectileGraphCmd.Flags().StringVarP(&gravity, "gravity", "g", "0,-0.1,0", "the gravity vector")
	projectileGraphCmd.Flags().StringVarP(&wind, "wind", "w", "-0.01,0,0", "the wind vector")
	rootCmd.AddCommand(projectileCmd)
	rootCmd.AddCommand(projectileGraphCmd)
}

func tick(p projectile, env environment) projectile {
	position := p.position.Add(p.velocity)
	velocity := p.velocity.Add(env.gravity).Add(env.wind)
	return projectile{
		position: position,
		velocity: velocity,
	}
}
