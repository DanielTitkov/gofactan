package gofactan

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestCalculateBartlettSphericity(t *testing.T) {
	data, err := CSVToMatrix("test/data/test01.csv", true)
	if err != nil {
		t.Errorf("failed to read data: %s", err)
	}

	expS := 14185.002857
	expP := 0.0

	s, p := CalculateBartlettSphericity(data)
	s = roundToPlace(s, 6)
	if s != expS {
		t.Errorf("expected statistic of %f but got %f", expS, s)
	}
	if p != expP {
		t.Errorf("expected p-value of %f but got %f", expP, p)
	}
}

func TestCalculateKMO(t *testing.T) {
	data, err := CSVToMatrix("test/data/test02.csv", true)
	if err != nil {
		t.Errorf("failed to read data: %s", err)
	}

	expectedItemsData := []float64{0.405516, 0.560049, 0.700033,
		0.705446, 0.829063, 0.848425,
		0.863502, 0.841143, 0.877076,
		0.839272}
	expectedItems := mat.NewVecDense(len(expectedItemsData), expectedItemsData)
	expectedTotal := 0.814985

	kmoItems, kmoTotal, err := CalculateKMO(data)
	if err != nil {
		t.Errorf("failed to calculate KMO: %s", err)
	}

	applyToVector(func(i int, v float64) float64 {
		return roundToPlace(v, 6)
	}, kmoItems)
	kmoTotal = roundToPlace(kmoTotal, 6)

	if !reflect.DeepEqual(expectedItems, kmoItems) {
		t.Errorf("expected %v, got %v", expectedItems, kmoItems)
	}
	if expectedTotal != kmoTotal {
		t.Errorf("expected %f, got %f", expectedTotal, kmoTotal)
	}
}
