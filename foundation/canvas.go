package foundation

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"strings"
)

const MAX_COLOR = 255

type Canvas struct {
	Width  int
	Height int
	pixels []Color
}

func NewCanvas(x, y int) *Canvas {
	c := Canvas{
		Width:  x,
		Height: y,
	}
	c.pixels = make([]Color, c.Width*c.Height)
	return &c
}

func (c *Canvas) GetPixel(x, y int) *Color {
	return &c.pixels[x+c.Width*y]
}
func (c *Canvas) SetPixel(x, y int, color *Color) {
	if x < c.Width && y < c.Height {
		c.pixels[x+c.Width*y] = *color
	}
}

func (c *Canvas) ToPPM() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintln("P3"))
	fmt.Fprintf(&builder, "%d %d\n", c.Width, c.Height)
	fmt.Fprintf(&builder, "%d\n", MAX_COLOR)
	rowLength := 0
	col := 0
	for _, pixel := range c.pixels {
		col += 1
		for _, component := range pixel {
			var str string
			if rowLength == 0 {
				str = fmt.Sprintf("%d", scaleColor(component))
			} else {
				str = fmt.Sprintf(" %d", scaleColor(component))
			}
			rowLength += len(str)
			if rowLength > 70 {
				fmt.Fprintln(&builder)
				str = fmt.Sprintf("%d", scaleColor(component))
				rowLength = len(str)
			}
			builder.WriteString(str)
		}
		if col == c.Width {
			fmt.Fprintln(&builder)
			col = 0
			rowLength = 0
		}
	}
	fmt.Fprintln(&builder)
	return builder.String()
}

func scaleColor(color float64) int64 {
	if color < 0 {
		return 0
	}
	if color > 1 {
		return MAX_COLOR
	}

	return int64(math.Round(MAX_COLOR * color))
}

func (c *Canvas) SavePNG(filename string) {
	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))

	for y := 0; y < c.Height; y += 1 {
		for x := 0; x < c.Width; x += 1 {
			pixel := c.GetPixel(x, y)
			color := pixel.GetRGBA()

			img.Set(x, y, color)
		}
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	png.Encode(f, img)
}
