package main

import (
	"fmt"
	"io/ioutil"

	base "github.com/varigg/raytacer-challenge/foundation"
)

type environment struct {
	gravity *base.Tuple
	wind    *base.Tuple
}

type projectile struct {
	position *base.Tuple
	velocity *base.Tuple
}

func main() {
	fmt.Println("hello world")
	p := projectile{
		position: base.NewPoint(0, 1, 0),
		velocity: base.NewVector(1, 1, 0).Normalize().Times(11),
	}
	e := environment{
		gravity: base.NewVector(0, -0.1, 0),
		wind:    base.NewVector(-0.01, 0, 0),
	}
	width, height := 1000, 500
	canvas := base.NewCanvas(width, height)
	red := base.NewColor(1, 0, 0)
	fmt.Printf("x: %d, y: %d\n", int(p.position.Y()), int(p.position.X()))
	for p.position.Y() > 0 {
		p = tick(e, p)
		fmt.Printf("x: %d, y: %d\n", int(p.position.X()), height-int(p.position.Y()))
		canvas.SetPixel(int(p.position.X()), height-int(p.position.Y()), red)

	}
	ppm := canvas.ToPPM()
	ioutil.WriteFile("canvas.ppm", []byte(ppm), 0600) //nolint

}

func tick(env environment, proj projectile) projectile {
	position := proj.position.Plus(proj.velocity)
	velocity := proj.velocity.Plus(env.gravity).Plus(env.wind)
	return projectile{
		position: position,
		velocity: velocity,
	}

}
