package core

import (
	"image/color"
	"math"
)

type Color struct {
	red   float64
	green float64
	blue  float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{
		red:   r,
		green: g,
		blue:  b,
	}
}

func (c *Color) Red() float64 {
	return c.red
}

func (c *Color) Green() float64 {
	return c.green
}

func (c *Color) Blue() float64 {
	return c.blue
}

func (c1 *Color) Add(c2 *Color) *Color {
	return &Color{
		red:   c1.Red() + c2.Red(),
		green: c1.Green() + c2.Green(),
		blue:  c1.Blue() + c2.Blue(),
	}
}

func (c1 *Color) Subtract(c2 *Color) *Color {
	return &Color{
		red:   c1.Red() - c2.Red(),
		green: c1.Green() - c2.Green(),
		blue:  c1.Blue() - c2.Blue(),
	}
}

func (c1 *Color) Multiply(s float64) *Color {
	return &Color{
		red:   c1.Red() * s,
		green: c1.Green() * s,
		blue:  c1.Blue() * s,
	}
}

func (c1 *Color) HadamardProduct(c2 *Color) *Color {
	return &Color{
		red:   c1.Red() * c2.Red(),
		green: c1.Green() * c2.Green(),
		blue:  c1.Blue() * c2.Blue(),
	}
}

func (c1 *Color) Equals(c2 *Color) bool {
	return equals(c1.Red(), c2.Red()) &&
		equals(c1.Green(), c2.Green()) && equals(c1.Blue(), c2.Blue())
}

func (c1 *Color) ToRGBA(maxValues int) color.RGBA {
	rgb := make([]uint8, 4)
	rgb[0] = ConvertFloatToColorValue(c1.Red(), maxValues)
	rgb[1] = ConvertFloatToColorValue(c1.Green(), maxValues)
	rgb[2] = ConvertFloatToColorValue(c1.Blue(), maxValues)
	rgb[3] = 0xFF
	return color.RGBA{rgb[0], rgb[1], rgb[2], rgb[3]}
}

func ConvertFloatToColorValue(f float64, maxValue int) uint8 {
	// Ensure the float value is within the range [0, 1]
	if f < 0 {
		f = 0
	} else if f > 1.0 {
		f = 1.0
	}

	// Multiply the float by maxValue and convert to integer
	result := uint8(math.Round(f * float64(maxValue)))

	return result
}
