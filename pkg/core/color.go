package core

import (
	"image/color"
	"math"
)

type Color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) Color {
	return Color{R: r, G: g, B: b}
}

func (c Color) Add(o Color) Color {
	return Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

func (c Color) Subtract(o Color) Color {
	return Color{c.R - o.R, c.G - o.G, c.B - o.B}
}

func (c Color) Multiply(s float64) Color {
	return Color{c.R * s, c.G * s, c.B * s}
}

func (c Color) HadamardProduct(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}

func (c Color) Equals(o Color) bool {
	return equals(c.R, o.R) && equals(c.G, o.G) && equals(c.B, o.B)
}

func (c Color) ToRGBA(maxValue int) color.RGBA {
	return color.RGBA{
		R: ConvertFloatToColorValue(c.R, maxValue),
		G: ConvertFloatToColorValue(c.G, maxValue),
		B: ConvertFloatToColorValue(c.B, maxValue),
		A: 0xFF,
	}
}

func ConvertFloatToColorValue(f float64, maxValue int) uint8 {
	f = math.Max(0, math.Min(1, f))
	return uint8(math.Round(f * float64(maxValue)))
}
