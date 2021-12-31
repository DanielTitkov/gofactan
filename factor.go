package gofactan

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func CalculateKMO(x *mat.Dense) (kmoPerVariable *mat.VecDense, kmoTotal float64, err error) {

	// pair-wise correlations
	xCorr := CorrMartix(x)
	partCorr, err := PartialCorr(x)
	if err != nil {
		return nil, 0, err
	}

	// fill matrix diagonals with zeros
	fillDiag(xCorr, 0)
	fillDiag(partCorr, 0)

	// square all elements
	xCorr.Apply(func(i, j int, v float64) float64 { return math.Pow(v, 2) }, xCorr)
	partCorr.Apply(func(i, j int, v float64) float64 { return math.Pow(v, 2) }, partCorr)

	// calculate KMO per variable
	corrSumsVec := axisSum(xCorr, "rows")
	partSumsVec := axisSum(partCorr, "rows")

	div, kmoPerVariable := &mat.VecDense{}, &mat.VecDense{}
	div.AddVec(corrSumsVec, partSumsVec)
	kmoPerVariable.DivElemVec(corrSumsVec, div)

	// calculate overall KMO
	corrSumTotal := matrixSum(xCorr)
	partSumTotal := matrixSum(partCorr)
	kmoTotal = corrSumTotal / (corrSumTotal + partSumTotal)

	fmt.Println(kmoTotal)

	return kmoPerVariable, kmoTotal, nil
}
