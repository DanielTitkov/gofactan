package gofactan

import (
	"gonum.org/v1/gonum/mat"
)

type Factor struct {
	Cfg FactorConfig
}

type FactorConfig struct {
	NFactors     int
	Rotation     string
	Method       string
	UseSMC       bool
	IsCorrMatrix bool
}

func NewFactor() *Factor {

}

func (f *Factor) Fit(x mat.Matrix) error {
	return nil
}
