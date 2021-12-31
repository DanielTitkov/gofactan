package gofactan

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

func CSVToMatrix(path string, dropHeader bool) (*mat.Dense, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	if len(lines) < 2 {
		return nil, errors.New("data has less than 2 lines")
	}

	if dropHeader {
		lines = lines[1:]
	}

	var rows, cols int
	var data []float64

	rows = len(lines)

	for _, line := range lines {
		cols = len(line)
		for _, val := range line {
			v, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}

			data = append(data, v)
		}
	}

	return mat.NewDense(rows, cols, data), nil
}
