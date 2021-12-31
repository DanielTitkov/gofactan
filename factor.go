package gofactan

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func CalculateKMO(x *mat.Dense) (kmoPerVariable []float64, kmoTotal float64, err error) {

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

	corrSums := axisSum(xCorr, "rows")
	partSums := axisSum(partCorr, "rows")

	// TODO: maybe use vectors in sums calculation
	corrSumsVec := mat.NewVecDense(len(corrSums), corrSums)
	partSumsVec := mat.NewVecDense(len(partSums), partSums)

	div, varKMOVec := &mat.VecDense{}, &mat.VecDense{}
	div.AddVec(corrSumsVec, partSumsVec)
	varKMOVec.DivElemVec(corrSumsVec, div)

	fmt.Println(partSums)
	fmt.Println(corrSums)
	fmt.Println(varKMOVec)

	// calculate overall KMO
	// corrSumsTotal := axisSum(xCorr, "cols")
	// partSumsTotal := axisSum(partCorr, "cols")

	return nil, 0, nil
}
