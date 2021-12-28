package gofactan

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestCorrMatrix(t *testing.T) {
	data := mat.NewDense(3, 4, []float64{
		1, 2, 3, 4,
		5, 4, 3, 2,
		1, 2, 1, 2,
	})

	exp := mat.NewDense(4, 4, []float64{
		1, 1, 0.4999999999999999, -0.4999999999999999,
		1, 1, 0.5000000000000001, -0.5000000000000001,
		0.4999999999999999, 0.5000000000000001, 1, 0.5000000000000001,
		-0.4999999999999999, -0.5000000000000001, 0.5000000000000001, 1,
	})

	res := CorrMartix(data)

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("expected %+v but got %+v", exp, res)
	}
}

func TestCovMatrix(t *testing.T) {
	data := mat.NewDense(3, 4, []float64{
		1, 2, 3, 4,
		5, 4, 3, 2,
		1, 2, 1, 2,
	})

	exp := mat.NewDense(4, 4, []float64{
		5.333333333333334, 2.666666666666667, 1.3333333333333333, -1.3333333333333333,
		2.666666666666667, 1.3333333333333333, 0.6666666666666667, -0.6666666666666667,
		1.3333333333333333, 0.6666666666666667, 1.3333333333333333, 0.6666666666666667,
		-1.3333333333333333, -0.6666666666666667, 0.6666666666666667, 1.3333333333333333,
	})

	res := CovMartix(data)

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("expected %+v but got %+v", exp, res)
	}
}

func TestCov2Corr(t *testing.T) {
	data := mat.NewDense(3, 4, []float64{ // TODO: add more data
		1, 2, 3, 4,
		5, 4, 3, 2,
		1, 2, 1, 2,
	})

	cov := CovMartix(data)
	expCorr := CorrMartix(data)
	expCorr.Apply(func(i, j int, v float64) float64 {
		return roundToPlace(v, 15)
	}, expCorr)

	corr, err := Cov2Corr(cov)
	if err != nil {
		t.Error(err)
	}

	corr.Apply(func(i, j int, v float64) float64 {
		return roundToPlace(v, 15)
	}, corr)

	if !reflect.DeepEqual(expCorr, corr) {
		t.Errorf("expected %+v but got %+v", expCorr, corr)
	}
}

func TestPatialCorr(t *testing.T) {
	data := mat.NewDense(4, 3, []float64{ // TODO: add more data
		12, 14, 15,
		24, 12, 52,
		35, 12, 41,
		23, 12, 42,
	})

	exp := mat.NewDense(3, 3, []float64{
		1.0, -0.730955, -0.50616,
		-0.730955, 1.0, -0.928701,
		-0.50616, -0.928701, 1.0,
	})

	res, err := PartialCorr(data)
	if err != nil {
		t.Error(err)
	}

	res.Apply(func(i, j int, v float64) float64 {
		return roundToPlace(v, 6)
	}, res)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %+v but got %+v", exp, res)
	}
}

func TestFloatSlice(t *testing.T) {
	exp := []float64{5, 5, 5, 5}
	res := floatSlice(4, 5)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %+v but got %+v", exp, res)
	}
}

func TestRoundToPlace(t *testing.T) {
	type testcase struct {
		value float64
		place int
		exp   float64
	}

	tcs := []testcase{
		{
			0.45343,
			3,
			0.453,
		},
		{
			0.05,
			1,
			0.1,
		},
		{
			0.5555555555,
			5,
			0.55556,
		},
		{
			0.4999999993,
			5,
			0.5,
		},
	}

	for i, tc := range tcs {
		if res := roundToPlace(tc.value, tc.place); res != tc.exp {
			t.Errorf("expected %f, but got %f in test case %d", tc.exp, res, i)
		}
	}
}
