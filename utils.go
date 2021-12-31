package gofactan

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

func CorrMartix(x mat.Matrix) *mat.Dense {
	res := &mat.SymDense{}
	stat.CorrelationMatrix(res, x, nil)
	return mat.DenseCopyOf(res)
}

func CovMartix(x mat.Matrix) *mat.Dense {
	res := &mat.SymDense{}
	stat.CovarianceMatrix(res, x, nil)
	return mat.DenseCopyOf(res)
}

func Cov2Corr(x mat.Matrix) (*mat.Dense, error) {
	rows, cols := x.Dims()
	if rows != cols {
		return nil, errors.New("square matrix required")
	}

	diag := &mat.DiagDense{}
	diag.DiagFrom(x)

	is := &mat.Dense{}
	ones := mat.NewDiagDense(rows, floatSlice(rows, 1))

	is.DivElem(ones, diag)

	is.Apply(func(i, j int, v float64) float64 { return math.Sqrt(v) }, is)

	diag.DiagFrom(is)
	repeatedDiagCols := repeatDiag(diag, "columns")
	repeatedDiagRows := repeatDiag(diag, "rows")

	res := &mat.Dense{}
	res.MulElem(repeatedDiagCols, x)
	res.MulElem(res, repeatedDiagRows)

	fillDiag(res, 1.0)

	return res, nil
}

func PartialCorr(x mat.Matrix) (*mat.Dense, error) {
	rows, cols := x.Dims()

	xCov := CovMartix(x)

	empty := mat.NewDense(cols, rows, nil)
	icvx := &mat.Dense{}

	if cols > rows {
		icvx = empty
	} else {
		icvx.Inverse(xCov)
	}

	icvxRows, icvxCols := icvx.Dims()
	ones := mat.NewDense(icvxRows, icvxCols, nil)
	fill(ones, -1)

	pCorr := &mat.Dense{}
	corr, err := Cov2Corr(icvx)
	if err != nil {
		return nil, err
	}
	pCorr.MulElem(ones, corr)

	fillDiag(pCorr, 1.0)

	return pCorr, nil
}

func floatSlice(n int, v float64) []float64 {
	var res []float64
	for i := 0; i < n; i++ {
		res = append(res, v)
	}

	return res
}

func repeatDiag(dm mat.Diagonal, direction string) *mat.Dense {
	// TODO is there a way to make this less ugly?
	r, c := dm.Dims()
	if r != c {
		panic(errors.New("matrix must be square"))
	}

	var diagData []float64
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			if i == j {
				diagData = append(diagData, dm.At(i, j))
			}
		}
	}

	res := mat.NewDense(r, r, nil)
	switch direction {
	case "rows":
		for i := 0; i < r; i++ {
			res.SetRow(i, diagData)
		}
	case "columns":
		for i := 0; i < r; i++ {
			res.SetCol(i, diagData)
		}
	default:
		panic(errors.New("unknown direction, must be 'rows' or 'columns'"))
	}

	return res
}

func fillDiag(m mat.Mutable, value float64) {
	r, c := m.Dims()
	if r != c {
		panic(errors.New("matrix must be square"))
	}

	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			if i == j {
				m.Set(i, j, value)
			}
		}
	}
}

func fill(m mat.Mutable, value float64) {
	r, c := m.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m.Set(i, j, value)
		}
	}
}

func roundToPlace(v float64, n int) float64 {
	buf := math.Pow10(n)
	return math.Round(v*buf) / buf
}

func axisSum(m mat.Matrix, axis string) mat.Vector {
	var dim int
	r, c := m.Dims()
	switch axis {
	case "rows":
		dim = r
	case "cols":
		dim = c
	default:
		panic(errors.New("unknown axis, must be rows or cols"))
	}
	sums := make([]float64, dim)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			switch axis {
			case "rows":
				sums[i] += m.At(i, j)
			case "cols":
				sums[j] += m.At(i, j)
			}
		}
	}

	return mat.NewVecDense(len(sums), sums)
}

func matrixSum(m mat.Matrix) float64 {
	r, c := m.Dims()
	var sum float64
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			sum += m.At(i, j)
		}
	}

	return sum
}

func applyToVector(fn func(i int, v float64) float64, vec *mat.VecDense) {
	c := vec.Len()
	for i := 0; i < c; i++ {
		val := vec.AtVec(i)
		newVal := fn(i, val)
		vec.SetVec(i, newVal)
	}
}
