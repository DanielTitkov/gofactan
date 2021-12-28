package gofactan

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func calculateKMO(x *mat.Dense) (kmoPerVariable []float64, kmoTotal float64, err error) {

	// pair-wise correlations
	xCorr := CorrMartix(x)

	fmt.Println(xCorr)
	return nil, 0, nil
}
