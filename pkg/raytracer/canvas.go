package raytracer

import (
	"fmt"
	"io"
)

const MAX_COLORS = 256

type Canvas struct {
	Height int
	Width  int
	pixels [][]Color
}

func NewCanvas(x, y int) *Canvas {
	p := make([][]Color, x)
	for i := range p {
		p[i] = make([]Color, y)
	}
	c := Canvas{
		Width:  x,
		Height: y,
		pixels: p,
	}
	return &c
}

func (c *Canvas) Get(x, y int) *Color {
	return &c.pixels[x][y]
}

func (c *Canvas) Set(x, y int, color *Color) {
	c.pixels[x][y] = *color
}

func (c *Canvas) Pixels() []*Color {
	pixels := make([]*Color, 0)
	for y := range c.Height - 1 {
		for x := range c.Width - 1 {
			pixels = append(pixels, c.Get(x, y))
		}
	}
	return pixels
}

func (c *Canvas) ToPPM(writer io.Writer) error {
	ppm := fmt.Sprintf("P3\n%d %d\n%d\n", c.Width, c.Height, MAX_COLORS-1)
	_, err := writer.Write([]byte(ppm))
	if err != nil {
		return err
	}
	var currentLineLength int

	for y := range c.Height {
		for x := range c.Width {
			rgb := c.Get(x, y).ToRGB(MAX_COLORS - 1)
			for color := range rgb {
				str := fmt.Sprintf("%d", rgb[color])
				if currentLineLength > 0 && currentLineLength+1+len(str) > 70 {
					// Add a new line if the accumulated line length exceeds 70 characters
					writer.Write([]byte("\n"))
					currentLineLength = 0
				}
				if currentLineLength > 0 {
					writer.Write([]byte(" "))
					currentLineLength++
				}
				writer.Write([]byte(str))
				currentLineLength += len(str)
			}
		}
		writer.Write([]byte("\n"))
		currentLineLength = 0
	}
	return nil
}
