package gofactan

import (
	"reflect"
	"testing"
)

func TestCalculateKMO(t *testing.T) {
	data, err := CSVToMatrix("test/data/test02.csv", true)
	if err != nil {
		t.Errorf("failed to read data: %s", err)
	}

	expectedItems := []float64{0.405516, 0.560049, 0.700033,
		0.705446, 0.829063, 0.848425,
		0.863502, 0.841143, 0.877076,
		0.839272}
	expectedTotal := 0.81498469767761361

	kmoItems, kmoTotal, err := CalculateKMO(data)
	if err != nil {
		t.Errorf("failed to calculate KMO: %s", err)
	}
	if !reflect.DeepEqual(expectedItems, kmoItems) {
		t.Errorf("expected %v, got %v", expectedItems, kmoItems)
	}
	if expectedTotal != kmoTotal {
		t.Errorf("expected %f, got %f", expectedTotal, kmoTotal)
	}
}
