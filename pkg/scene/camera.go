package scene

import (
	"math"
	"sync"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
)

type Camera struct {
	HSize      int
	VSize      int
	FOV        float64
	PixelSize  float64
	HalfWidth  float64
	HalfHeight float64

	transform core.Matrix
	// inverse and origin cache transform.Invert() and the camera position;
	// RayForPixel reads them for every pixel.
	inverse core.Matrix
	origin  core.Tuple
}

func (c *Camera) SetTransform(m core.Matrix) {
	c.transform = m
	c.inverse = m.Invert()
	c.origin = c.inverse.MultiplyWithTuple(core.NewPoint(0, 0, 0))
}

func (c *Camera) Transform() core.Matrix { return c.transform }

func ViewTransform(from, to, up core.Tuple) core.Matrix {
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
		HSize: hsize,
		VSize: vsize,
		FOV:   fov,
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
	c.SetTransform(core.Identity(4))

	return c
}

func (c *Camera) RayForPixel(x, y int) *objects.Ray {
	xOffset := (float64(x) + .5) * c.PixelSize
	yOffset := (float64(y) + .5) * c.PixelSize
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	pixel := c.inverse.MultiplyWithTuple(core.NewPoint(worldX, worldY, -1))
	direction := pixel.Subtract(c.origin).Normalize()
	return objects.NewRay(c.origin, direction)
}

func (c *Camera) Render(w *World) *core.Canvas {
	image := core.NewCanvas(c.HSize, c.VSize)
	var wg sync.WaitGroup
	for y := 0; y < c.VSize; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := 0; x < c.HSize; x++ {
				image.Set(x, y, w.ColorAt(c.RayForPixel(x, y)))
			}
		}(y)
	}
	wg.Wait()
	return image
}
