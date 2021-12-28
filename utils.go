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

func roundToPlace(v float64, n int) float64 {
	buf := math.Pow10(n)
	return math.Round(v*buf) / buf
}
