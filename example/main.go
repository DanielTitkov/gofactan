package main

import (
	"fmt"

	fa "github.com/DanielTitkov/gofactan"
	"gonum.org/v1/gonum/mat"
)

func main() {
	fmt.Println("Run examples")
	data := mat.NewDense(4, 3, []float64{ // TODO: add more data
		12, 14, 15,
		24, 12, 52,
		35, 12, 41,
		23, 12, 42,
	})

	// exp := mat.NewDense(3, 3, []float64{
	// 	1.0, -0.730955, -0.50616,
	// 	-0.730955, 1.0, -0.928701,
	// 	-0.50616, -0.928701, 1.0,
	// })

	res, err := fa.PartialCorr(data)
	fmt.Println(res, err)
	// if err != nil {
	// 	t.Error(err)
	// }

	// if !reflect.DeepEqual(exp, res) {
	// 	t.Errorf("expected %+v but got %+v", exp, res)
	// }
}
