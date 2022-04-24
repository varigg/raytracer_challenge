package foundation

import (
	"image/color"
	"math"
)

type Color [3]float64

func (c *Color) Red() float64 {
	return c[0]
}
func (c *Color) Green() float64 {
	return c[1]
}
func (c *Color) Blue() float64 {
	return c[2]
}

func (c *Color) Plus(c1 *Color) *Color {
	return NewColor(c.Red()+c1.Red(), c.Green()+c1.Green(), c.Blue()+c1.Blue())
}

func (c *Color) Minus(c1 *Color) *Color {
	return NewColor(c.Red()-c1.Red(), c.Green()-c1.Green(), c.Blue()-c1.Blue())
}

func (c *Color) Blend(c1 *Color) *Color {
	return NewColor(c.Red()*c1.Red(), c.Green()*c1.Green(), c.Blue()*c1.Blue())
}
func (c *Color) Times(s float64) *Color {
	return NewColor(c.Red()*s, c.Green()*s, c.Blue()*s)
}

func (c1 *Color) Equals(c2 *Color) bool {
	return FromColor(c1).Equals(FromColor(c2))
}
func NewColor(r, g, b float64) *Color {
	c := Color{r, g, b}
	return &c
}
func (c Color) GetRGBA() color.Color {
	r := uint8(math.Min(math.Max(math.Round(c.Red()*255), 0), 255))
	g := uint8(math.Min(math.Max(math.Round(c.Green()*255), 0), 255))
	b := uint8(math.Min(math.Max(math.Round(c.Blue()*255), 0), 255))

	return color.RGBA{r, g, b, 0xFF}
}
