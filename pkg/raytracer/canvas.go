package raytracer

import (
	"bytes"
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
	fmt.Printf("setting %d, %d to %f,%f,%f\n", x, y, color.red, color.blue, color.green)
	c.pixels[x][y] = *color
}

func (c *Canvas) ToPPM(writer io.Writer) (int, error) {
	ppm := fmt.Sprintf("P3\n%d %d\n%d\n", c.Width, c.Height, MAX_COLORS-1)
	length := 0
	n, err := writer.Write([]byte(ppm))
	length += n
	if err != nil {
		return length, err
	}
	for y := range len(c.pixels[0]) {
		line := RowToString(y, c.pixels)
		//fmt.Println(line)
		n, err = writer.Write([]byte(line))
		length += n
		if err != nil {
			return length, err
		}
	}
	return length, nil
}
func RowToString(row int, pixels [][]Color) []byte {
	convertedRow := make([]string, 0)
	for x := range len(pixels) {
		r, g, b := pixels[x][row].ToRGB(MAX_COLORS - 1)
		convertedRow = append(convertedRow, fmt.Sprintf("%d", r))
		convertedRow = append(convertedRow, fmt.Sprintf("%d", g))
		convertedRow = append(convertedRow, fmt.Sprintf("%d", b))
	}
	var buf bytes.Buffer
	var currentLineLength int
	for i := range convertedRow {
		// max length of 70 characters per line
		if currentLineLength > 0 && currentLineLength+1+len(convertedRow[i]) > 70 {
			// Add a new line if the accumulated line length exceeds maxLineLength
			buf.WriteString("\n")
			currentLineLength = 0 // Reset current line length after line break
		}
		if currentLineLength > 0 {
			buf.WriteString(" ")
			currentLineLength++
		}
		buf.WriteString(convertedRow[i])
		currentLineLength += len(convertedRow[i])
		// add space after each value except at end of row
	}
	buf.WriteString("\n")
	return buf.Bytes()
}
