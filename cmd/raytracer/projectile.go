package raytracer

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/varigg/raytracer-challenge/pkg/raytracer"
)

var wind string
var gravity string
var velocity string
var origin string

var projectileCmd = &cobra.Command{
	Use:     "projectile",
	Aliases: []string{"chapter1"},
	Short:   "computes how far a projectile flies until it hits y=0",
	Run: func(cmd *cobra.Command, args []string) {
		p := projectile{
			position: raytracer.NewPointFromString(origin),
			velocity: raytracer.NewVectorFromString(velocity),
		}
		env := environment{
			gravity: raytracer.NewVectorFromString(gravity),
			wind:    raytracer.NewVectorFromString(wind),
		}
		distance := 0
		for p.position.Y() > 0 {
			fmt.Println(p.position.X())
			p = tick(p, env)
			distance += 1
		}
		fmt.Println(distance)

	},
}

type projectile struct {
	position raytracer.Tuple
	velocity raytracer.Tuple
}

type environment struct {
	gravity raytracer.Tuple
	wind    raytracer.Tuple
}

func init() {
	projectileCmd.Flags().StringVarP(&origin, "origin", "o", "0,1,0", "the starting point")
	projectileCmd.Flags().StringVarP(&velocity, "velocity", "v", "1,1,0", "the initial velocity")
	projectileCmd.Flags().StringVarP(&gravity, "gravity", "g", "0,-0.1,0", "the gravity vector")
	projectileCmd.Flags().StringVarP(&wind, "wind", "w", "-0.01,0,0", "the wind vector")
	rootCmd.AddCommand(projectileCmd)
}

func tick(p projectile, env environment) projectile {
	position := p.position.Add(p.velocity)
	velocity := p.velocity.Add(env.gravity).Add(env.wind)
	return projectile{
		position: position,
		velocity: velocity,
	}
}
