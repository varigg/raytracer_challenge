package main

import "github.com/varigg/raytacer-challenge/foundation"

func main() {
	// Set up canvas
	pixels := 500
	canvas := foundation.NewCanvas(pixels, pixels)

	// Ray starting at Z -5
	rayOrigin := foundation.NewPoint(0, 0, -5)

	// Wall at Z 10
	wallZ := 10.0
	wallSize := 7.0

	pixelSize := wallSize / float64(pixels)

	half := wallSize / 2

	shape := foundation.NewSphere()

	for y := 0; y < pixels; y += 1 {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < pixels; x += 1 {
			worldX := -half + pixelSize*float64(x)
			// the point on the wall we are shooting the ray at
			position := foundation.NewPoint(worldX, worldY, wallZ)
			// ray starting at origin with a vector towards the target point
			r := foundation.NewRay(rayOrigin, position.Minus(rayOrigin).Normalize())
			r.AddIntersections(shape.Intersect(r)...)
			if r.Hit() != nil {
				canvas.SetPixel(x, y, foundation.NewColor(1, 0, 0))
			}
		}
	}

	canvas.SavePNG("test.png")
}
