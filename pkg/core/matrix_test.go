package core_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varigg/raytracer-challenge/pkg/core"
)

func TestCreateMatrix2(t *testing.T) {
	m := *core.NewMatrix([][]float64{
		{-3, 5},
		{1, -2},
	})
	assert.Equal(t, float64(-3), m[0][0])
	assert.Equal(t, float64(5), m[0][1])
	assert.Equal(t, float64(1), m[1][0])
	assert.Equal(t, float64(-2), m[1][1])
}

func TestCreateMatrix3(t *testing.T) {
	m := *core.NewMatrix([][]float64{
		{-3, 5, 0},
		{1, -2, 7},
		{0, 1, 1},
	})
	assert.Equal(t, float64(-3), m[0][0])
	assert.Equal(t, float64(-2), m[1][1])
	assert.Equal(t, float64(1), m[2][2])

}
func TestCreateMatrix4(t *testing.T) {
	m := *core.NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	assert.Equal(t, float64(1), m[0][0])
	assert.Equal(t, float64(4), m[0][3])
	assert.Equal(t, float64(5.5), m[1][0])
	assert.Equal(t, float64(7.5), m[1][2])
	assert.Equal(t, float64(11), m[2][2])
	assert.Equal(t, float64(13.5), m[3][0])
	assert.Equal(t, float64(15.5), m[3][2])
}

func TestMatrixEquivalency(t *testing.T) {
	data := [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}
	m1 := core.NewMatrix(data)
	m2 := core.NewMatrix(data)
	assert.Equal(t, m1, m2)

	data = [][]float64{
		{1, 2, 3, 4},
	}
	m2 = core.NewMatrix(data)
	assert.NotEqual(t, m1, m2)

	data = [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	m2 = core.NewMatrix(data)
	assert.NotEqual(t, m1, m2)
	assert.NotEqual(t, nil, m2)
}

func TestMatrixMultiplication(t *testing.T) {
	data := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	m1 := core.NewMatrix(data)

	data = [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}
	m2 := core.NewMatrix(data)

	data = [][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}
	expected := core.NewMatrix(data)
	assert.Equal(t, expected, m1.Times(m2))

	m2 = core.Identity4()
	assert.Equal(t, m1, m1.Times(m2))
}

func TestMatrixMultiplicationByTyple(t *testing.T) {
	data := [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	m := core.NewMatrix(data)

	tuple := core.NewTuple(1, 2, 3, 1)

	expected := core.NewTuple(18, 24, 33, 1)
	assert.Equal(t, expected, m.MultiplyWithTuple(tuple))
	m = core.Identity4()
	assert.Equal(t, tuple, m.MultiplyWithTuple(tuple))
}

func TestTransposeMatrix(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{0.0, 9.0, 3.0, 0.0},
		{9.0, 8.0, 0.0, 8.0},
		{1.0, 8.0, 5.0, 3.0},
		{0.0, 0.0, 5.0, 8.0},
	})

	expected := core.NewMatrix([][]float64{
		{0.0, 9.0, 1.0, 0.0},
		{9.0, 8.0, 8.0, 0.0},
		{3.0, 0.0, 5.0, 5.0},
		{0.0, 8.0, 3.0, 8.0},
	})

	//assert.True(t, expected.Equals(result))
	assert.Equal(t, expected, m.Transpose())
	assert.Equal(t, core.Identity4(), core.Identity4().Transpose())
}

func TestInvert2x2Matrix(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{1.0, 5.0},
		{-3.0, 2.0},
	})

	//assert.True(t, expected.Equals(result))
	assert.Equal(t, float64(17), m.Determinant())

}

func TestSubMatrix(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})
	sub := core.NewMatrix([][]float64{
		{-3.0, 2.0},
		{0, 6.0},
	})

	assert.Equal(t, sub, m.SubMatrix(0, 2))

	m = core.NewMatrix([][]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})
	sub = core.NewMatrix([][]float64{
		{-6.0, 1.0, 6},
		{-8, 8, 6.0},
		{-7, -1, 1},
	})

	assert.Equal(t, sub, m.SubMatrix(2, 1))
}

func TestDeterminant(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{1, 5},
		{-3, 2},
	})
	assert.Equal(t, float64(17), m.Determinant())
}
func TestSubMinors(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	assert.Equal(t, float64(25), m.Minor(1, 0))
}

func TestCofactor(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	assert.Equal(t, float64(-12), m.Minor(0, 0))
	assert.Equal(t, float64(-12), m.Cofactor(0, 0))
	assert.Equal(t, float64(25), m.Minor(1, 0))
	assert.Equal(t, float64(-25), m.Cofactor(1, 0))
}

func TestDeterminants3x3(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})
	assert.Equal(t, float64(56), m.Cofactor(0, 0))
	assert.Equal(t, float64(12), m.Cofactor(0, 1))
	assert.Equal(t, float64(-46), m.Cofactor(0, 2))
	assert.Equal(t, float64(-196), m.Determinant())
}
func TestDeterminants4x4(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})
	assert.Equal(t, float64(690), m.Cofactor(0, 0))
	assert.Equal(t, float64(447), m.Cofactor(0, 1))
	assert.Equal(t, float64(210), m.Cofactor(0, 2))
	assert.Equal(t, float64(51), m.Cofactor(0, 3))
	assert.Equal(t, float64(-4071), m.Determinant())
}

func TestInvertible(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	})
	assert.Equal(t, float64(-2120), m.Determinant())
	assert.True(t, m.IsInvertible())
	m = core.NewMatrix([][]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -4, 1, -5},
		{0, 0, 0, 0},
	})
	assert.False(t, m.IsInvertible())
}

func TestInversion(t *testing.T) {
	m := core.NewMatrix([][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	})

	m1 := m.Invert()
	assert.Equal(t, float64(532), m.Determinant())
	assert.Equal(t, float64(-160), m.Cofactor(2, 3))
	assert.Equal(t, float64(105), m.Cofactor(3, 2))
	assert.Equal(t, float64(-160.0/532.0), (*m1)[3][2])

	expected := core.NewMatrix([][]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	})
	t.Log(m1)
	assert.True(t, m1.Equals(expected))

	m1 = core.NewMatrix([][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})

	expected = core.NewMatrix([][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	})

	assert.True(t, m1.Invert().Equals(expected))

	m2 := m.Times(m1)

	assert.True(t, m2.Times(m1.Invert()).Equals(m))
	assert.True(t, core.Identity(4).Equals(core.Identity(4).Invert()))
}
func TestTranslation(t *testing.T) {
	transform := core.TranslationMatrix(5, -3, 2)
	p := core.NewPoint(-3, 4, 5)
	assert.Equal(t, core.NewPoint(2, 1, 7), transform.MultiplyWithTuple(p))
	assert.Equal(t, core.NewPoint(-8, 7, 3), transform.Invert().MultiplyWithTuple(p))
	v := core.NewVector(-3, 4, 5)
	assert.Equal(t, v, transform.MultiplyWithTuple(v))
}

func TestScaling(t *testing.T) {
	transform := core.ScalingMatrix(2, 3, 4)
	p := core.NewPoint(4, 6, 8)
	assert.Equal(t, core.NewPoint(8, 18, 32), transform.MultiplyWithTuple(p))
	v := core.NewVector(4, 6, 8)
	assert.Equal(t, core.NewVector(8, 18, 32), transform.MultiplyWithTuple(v))
	assert.Equal(t, core.NewVector(2, 2, 2), transform.Invert().MultiplyWithTuple(v))
	transform = core.ScalingMatrix(-1, 1, 1)
	assert.Equal(t, core.NewPoint(-4, 6, 8), transform.MultiplyWithTuple(p))
}

func TestReflection(t *testing.T) {
	transform := core.ScalingMatrix(-1, 1, 1)
	p := core.NewPoint(2, 3, 4)
	assert.Equal(t, core.NewPoint(-2, 3, 4), transform.MultiplyWithTuple(p))
}

func TestRotateX(t *testing.T) {
	//Scenario​: Rotating a point around the x axis
	p := core.NewPoint(0, 1, 0)
	half_quarter := core.RotationMatrixX(math.Pi / 4)
	full_quarter := core.RotationMatrixX(math.Pi / 2)
	assert.True(t, core.NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2).Equals(half_quarter.MultiplyWithTuple(p)))
	assert.True(t, core.NewPoint(0, 0, 1).Equals(full_quarter.MultiplyWithTuple(p)))
	assert.True(t, full_quarter.MultiplyWithTuple(p).Equals(half_quarter.MultiplyWithTuple(half_quarter.MultiplyWithTuple(p))))
}

func TestRotateY(t *testing.T) {
	//Scenario​: Rotating a point around the x axis
	p := core.NewPoint(0, 0, 1)
	half_quarter := core.RotationMatrixY(math.Pi / 4)
	full_quarter := core.RotationMatrixY(math.Pi / 2)
	assert.True(t, core.NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2).Equals(half_quarter.MultiplyWithTuple(p)))
	t.Log(half_quarter.MultiplyWithTuple(p))
	assert.True(t, core.NewPoint(1, 0, 0).Equals(full_quarter.MultiplyWithTuple(p)))
	t.Log(full_quarter.MultiplyWithTuple(p))
}

func TestRotateZ(t *testing.T) {
	//Scenario​: Rotating a point around the x axis
	p := core.NewPoint(0, 1, 0)
	half_quarter := core.RotationMatrixZ(math.Pi / 4)
	full_quarter := core.RotationMatrixZ(math.Pi / 2)
	assert.True(t, core.NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0).Equals(half_quarter.MultiplyWithTuple(p)))
	assert.True(t, core.NewPoint(-1, 0, 0).Equals(full_quarter.MultiplyWithTuple(p)))
}

func TestShearing(t *testing.T) {
	//Scenario​: Rotating a point around the x axis
	p := core.NewPoint(2, 3, 4)
	transform := core.Shearing(1, 0, 0, 0, 0, 0)
	t.Log(transform.MultiplyWithTuple(p))
	assert.True(t, core.NewPoint(5, 3, 4).Equals(transform.MultiplyWithTuple(p)))
	transform = core.Shearing(0, 1, 0, 0, 0, 0)
	t.Log(transform.MultiplyWithTuple(p))
	assert.True(t, core.NewPoint(6, 3, 4).Equals(transform.MultiplyWithTuple(p)))
	transform = core.Shearing(0, 0, 1, 0, 0, 0)
	t.Log(transform.MultiplyWithTuple(p))
	assert.True(t, core.NewPoint(2, 5, 4).Equals(transform.MultiplyWithTuple(p)))
	transform = core.Shearing(0, 0, 0, 1, 0, 0)
	t.Log(transform.MultiplyWithTuple(p))
	assert.True(t, core.NewPoint(2, 7, 4).Equals(transform.MultiplyWithTuple(p)))
	transform = core.Shearing(0, 0, 0, 0, 1, 0)
	t.Log(transform.MultiplyWithTuple(p))
	assert.True(t, core.NewPoint(2, 3, 6).Equals(transform.MultiplyWithTuple(p)))
	transform = core.Shearing(0, 0, 0, 0, 0, 1)
	t.Log(transform.MultiplyWithTuple(p))
	assert.True(t, core.NewPoint(2, 3, 7).Equals(transform.MultiplyWithTuple(p)))

}

func TestChaining(t *testing.T) {
	p := core.NewPoint(1, 0, 1)
	a := core.RotationMatrixX(math.Pi / 2)
	b := core.ScalingMatrix(5, 5, 5)
	c := core.TranslationMatrix(10, 5, 7)
	p2 := a.MultiplyWithTuple(p)
	assert.True(t, core.NewPoint(1, -1, 0).Equals(p2))
	p3 := b.MultiplyWithTuple(p2)
	assert.True(t, core.NewPoint(5, -5, 0).Equals(p3))
	p4 := c.MultiplyWithTuple(p3)
	assert.True(t, core.NewPoint(15, 0, 7).Equals(p4))
	assert.True(t, p4.Equals(c.Times(b).Times(a).MultiplyWithTuple(p)))
	transform := core.Identity(4).RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7)
	assert.True(t, p4.Equals(transform.MultiplyWithTuple(p)))
	assert.True(t, p4.Equals(p.Transform(transform)))
}
