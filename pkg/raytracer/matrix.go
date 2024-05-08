package raytracer

import (
	"fmt"
	"math"
)

// Matrix is a matrix of floating point numbers.
type Matrix [][]float64

// NewMatrix returns a new Matrix object.
func NewMatrix(data [][]float64) *Matrix {
	if len(data) != len(data[0]) {
		return nil
	}
	m := Matrix(data)
	return &m
}

func NewEmptyMatrix(size int) *Matrix {
	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		row := make([]float64, size)
		for j := 0; j < size; j++ {
			row[j] = 0.0
		}
		result[i] = row
	}
	return NewMatrix(result)
}

func (m *Matrix) Size() int {
	return len(*m)
}

// Identify4 returns an identity matrix of size 4

func Identity4() *Matrix {
	data := [][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	m := Matrix(data)
	return &m
}

func Identity(size int) *Matrix {
	m := NewEmptyMatrix(size)
	for j := 0; j < size; j++ {
		(*m)[j][j] = 1.0
	}
	return m
}

func (m1 *Matrix) Equals(m2 *Matrix) bool {
	if (m1 == nil) != (m2 == nil) {
		return false
	}
	if len(*m1) != len(*m2) {
		return false
	}
	for i, row := range *m1 {
		for j := range row {
			if !equals((*m1)[i][j], (*m2)[i][j]) {
				return false
			}
		}
	}

	return true
}

func (m1 *Matrix) Times(m2 *Matrix) *Matrix {
	size := len(*m1)
	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		row := make([]float64, size)
		for j := 0; j < size; j++ {
			iTuple := NewTuple((*m1)[i][0], (*m1)[i][1], (*m1)[i][2], (*m1)[i][3])
			jTuple := NewTuple((*m2)[0][j], (*m2)[1][j], (*m2)[2][j], (*m2)[3][j])
			row[j] = iTuple.Dot(jTuple)
		}
		result[i] = row
	}
	return NewMatrix(result)
}

func (m *Matrix) MultiplyWithTuple(t *Tuple) *Tuple {
	size := len(*m)
	result := make([]float64, size)
	for i := 0; i < size; i++ {
		iTuple := NewTuple((*m)[i][0], (*m)[i][1], (*m)[i][2], (*m)[i][3])
		result[i] = iTuple.Dot(t)
	}
	return NewTuple(result[0], result[1], result[2], result[3])
}

func (m *Matrix) Transpose() *Matrix {
	col := 0
	for i := 0; i < m.Size(); i++ {
		for j := col; j < m.Size(); j++ {
			temp := (*m)[i][j]
			(*m)[i][j] = (*m)[j][i]
			(*m)[j][i] = temp
		}
		col++
	}
	return m
}

func (m *Matrix) Determinant() float64 {
	data := *m
	var determinant float64
	if len(data) == 2 {
		determinant = data[0][0]*data[1][1] - data[0][1]*data[1][0]
	} else {
		for column := 0; column < len(data); column += 1 {
			determinant += data[0][column] * m.Cofactor(0, column)

		}
	}

	return determinant

}

func (m *Matrix) SubMatrix(rowIndex, columnIndex int) *Matrix {
	data := *m
	result := make([][]float64, 0)
	for i := 0; i < m.Size(); i++ {
		if i != rowIndex {
			row := make([]float64, 0)
			for j := 0; j < m.Size(); j++ {
				if j != columnIndex {
					row = append(row, data[i][j])
				}
			}
			result = append(result, row)
		}

	}
	return NewMatrix(result)
}

func (m *Matrix) Minor(rowIndex, columnIndex int) float64 {
	return m.SubMatrix(rowIndex, columnIndex).Determinant()
}

func (m *Matrix) Cofactor(rowIndex, columnIndex int) float64 {
	var sign float64
	if (rowIndex+columnIndex)%2 == 0 {
		sign = 1.0
	} else {
		sign = -1.0
	}

	return sign * m.SubMatrix(rowIndex, columnIndex).Determinant()
}
func (m *Matrix) IsInvertible() bool {
	return m.Determinant() != 0
}
func (m *Matrix) Invert() *Matrix {
	size := len(*m)
	result := NewEmptyMatrix(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c := m.Cofactor(i, j)
			(*result)[j][i] = c / m.Determinant()
		}
	}
	return result
}

func TranslationMatrix(x, y, z float64) *Matrix {
	result := NewEmptyMatrix(4)

	(*result)[0][0] = 1
	(*result)[0][3] = x
	(*result)[1][1] = 1
	(*result)[1][3] = y
	(*result)[2][2] = 1
	(*result)[2][3] = z
	(*result)[3][3] = 1
	return result
}

func (m *Matrix) Translate(x, y, z float64) *Matrix {
	return TranslationMatrix(x, y, z).Times(m)
}

func ScalingMatrix(x, y, z float64) *Matrix {
	result := NewEmptyMatrix(4)
	(*result)[0][0] = x
	(*result)[1][1] = y
	(*result)[2][2] = z
	(*result)[3][3] = 1
	return result
}

func (m *Matrix) Scale(x, y, z float64) *Matrix {
	return ScalingMatrix(x, y, z).Times(m)
}
func RotationMatrixX(radians float64) *Matrix {
	result := NewEmptyMatrix(4)
	(*result)[0][0] = 1
	(*result)[1][1] = math.Cos(radians)
	(*result)[1][2] = -math.Sin(radians)
	(*result)[2][1] = math.Sin(radians)
	(*result)[2][2] = math.Cos(radians)
	(*result)[3][3] = 1
	return result
}

func RotationMatrixY(radians float64) *Matrix {
	result := NewEmptyMatrix(4)
	(*result)[0][0] = math.Cos(radians)
	(*result)[0][2] = math.Sin(radians)
	(*result)[1][1] = 1
	(*result)[2][0] = -math.Sin(radians)
	(*result)[2][2] = math.Cos(radians)
	(*result)[3][3] = 1
	return result
}

func RotationMatrixZ(radians float64) *Matrix {
	result := NewEmptyMatrix(4)
	(*result)[0][0] = math.Cos(radians)
	(*result)[0][1] = -math.Sin(radians)
	(*result)[1][0] = -math.Sin(radians)
	(*result)[1][1] = math.Cos(radians)
	(*result)[2][2] = 1
	(*result)[3][3] = 1
	return result
}

func (m *Matrix) RotateX(radians float64) *Matrix {
	return (RotationMatrixX(radians).Times(m))
}

func (m *Matrix) RotateY(radians float64) *Matrix {
	return RotationMatrixY(radians).Times(m)
}

func (m *Matrix) RotateZ(radians float64) *Matrix {
	return RotationMatrixZ(radians).Times(m)
}
func Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	result := NewEmptyMatrix(4)
	(*result)[0][0] = 1
	(*result)[0][1] = xy
	(*result)[0][2] = xz
	(*result)[1][0] = yx
	(*result)[1][1] = 1
	(*result)[1][2] = yz
	(*result)[2][0] = zx
	(*result)[2][1] = zy
	(*result)[2][2] = 1
	(*result)[3][3] = 1
	return result
}

func (m *Matrix) Print() {
	for i := 0; i < m.Size(); i++ {
		for j := 0; j < m.Size(); j++ {
			fmt.Print((*m)[i][j])
			if j < m.Size()-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
