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

	_, err := gofactan.CSVToMatrix("test/data/test02.csv", true)
	if err != nil {
		log.Fatalln(err)
	}
}
