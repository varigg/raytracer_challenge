package core

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strconv"
)

const maxColorValue = 255

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

func (c *Canvas) ToPPM(w io.Writer) error {
	bw := bufio.NewWriter(w)
	fmt.Fprintf(bw, "P3\n%d %d\n%d\n", c.Width, c.Height, maxColorValue)
	var lineLen int
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			rgb := c.Get(x, y).ToRGBA(maxColorValue)
			for _, v := range []uint8{rgb.R, rgb.G, rgb.B} {
				s := strconv.Itoa(int(v))
				if lineLen > 0 && lineLen+1+len(s) > 70 {
					bw.WriteByte('\n')
					lineLen = 0
				}
				if lineLen > 0 {
					bw.WriteByte(' ')
					lineLen++
				}
				bw.WriteString(s)
				lineLen += len(s)
			}
		}
		bw.WriteByte('\n')
		lineLen = 0
	}
	return bw.Flush()
}

func (c *Canvas) SavePPM(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("create %s: %w", fileName, err)
	}
	defer file.Close()
	return c.ToPPM(file)
}

func (c *Canvas) SavePNG(filename string) error {
	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			img.Set(x, y, c.Get(x, y).ToRGBA(maxColorValue))
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create %s: %w", filename, err)
	}
	defer f.Close()
	return png.Encode(f, img)
}
