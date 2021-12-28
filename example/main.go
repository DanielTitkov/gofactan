package main

import (
	"fmt"
	"reflect"

	fa "github.com/DanielTitkov/gofactan"
	"gonum.org/v1/gonum/mat"
)

func main() {
	fmt.Println("Run examples")
	data := mat.NewDense(3, 4, []float64{ // TODO: add more data
		1, 2, 3, 4,
		5, 4, 3, 2,
		1, 2, 1, 2,
	})

	cov := fa.CovMartix(data)
	expCorr := fa.CorrMartix(data)
	fmt.Println("cov", cov)
	fmt.Println("exp corr", expCorr)
	corr, err := fa.Cov2Corr(cov)
	if err != nil {
		fmt.Println(err)
	}

	if !reflect.DeepEqual(expCorr, corr) {
		fmt.Printf("expected %+v but got %+v", expCorr, corr)
	}
}
