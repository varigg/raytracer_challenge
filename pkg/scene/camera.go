package scene

import (
	"math"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
)

type Camera struct {
	HSize int
	VSize int
	FOV   float64
	// Todo: precompute inverse
	Transform  *core.Matrix
	PixelSize  float64
	HalfWidth  float64
	HalfHeight float64
}

func ViewTransform(from, to, up *core.Tuple) *core.Matrix {
	forward := to.Subtract(from).Normalize()
	left := forward.Cross(up.Normalize())
	trueUp := left.Cross(forward)
	orientation := core.NewMatrix([][]float64{
		{left.X, left.Y, left.Z, 0},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	})
	//return orientation.Translate(-from.X, -from.Y, -from.Z)
	return orientation.Times(core.TranslationMatrix(-from.X, -from.Y, -from.Z))
}

func NewCamera(hsize, vsize int, fov float64) *Camera {
	c := &Camera{
		HSize:     hsize,
		VSize:     vsize,
		FOV:       fov,
		Transform: core.Identity4(),
	}

	halfView := math.Tan(fov / 2)
	aspectRatio := float64(hsize) / float64(vsize)
	if aspectRatio >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspectRatio
	} else {
		c.HalfWidth = halfView * aspectRatio
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2) / float64(c.HSize)

	return c
}

func (c *Camera) RayForPixel(x, y int) *objects.Ray {
	xOffset := (float64(x) + .5) * c.PixelSize
	yOffset := (.5 + float64(y)) * c.PixelSize
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	// inverse can be precomputed
	transformI := c.Transform.Invert()
	pixel := transformI.MultiplyWithTuple(core.NewPoint(worldX, worldY, -1))
	// this can be precomputed
	origin := transformI.MultiplyWithTuple(core.NewPoint(0, 0, 0))

	direction := pixel.Subtract(origin).Normalize()

	r := objects.NewRay(origin, direction)
	return r
}

func (c *Camera) Render(w *World) *core.Canvas {
	image := core.NewCanvas(c.HSize, c.VSize)
	for y := range c.VSize {
		for x := range c.HSize {
			ray := c.RayForPixel(x, y)
			color := w.ColorAt(ray)
			image.Set(x, y, color)
		}
	}
	return image
}
