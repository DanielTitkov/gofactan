package main

import (
	"fmt"
	"log"

	"github.com/DanielTitkov/gofactan"
)

func main() {
	fmt.Println("Run examples")
	// data := mat.NewDense(4, 3, []float64{ // TODO: add more data
	// 	12, 14, 15,
	// 	24, 12, 52,
	// 	35, 12, 41,
	// 	23, 12, 42,
	// })

	m, err := gofactan.CSVToMatrix("test/data/test02.csv", true)
	if err != nil {
		log.Fatalln(err)
	}

	// expectedItems := []float64{0.405516, 0.560049, 0.700033,
	// 	0.705446, 0.829063, 0.848425,
	// 	0.863502, 0.841143, 0.877076,
	// 	0.839272}
	// expectedTotal := 0.81498469767761361

	kmoItems, kmoTotal, err := gofactan.CalculateKMO(m)

	fmt.Println(kmoItems, kmoTotal, err)
}
