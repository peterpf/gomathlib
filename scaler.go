package gomathlib

import (
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/floats"
)

type Scaler interface {
	Fit(data []float64)                 // Fit the scaler to the data
	Transform(v float64) float64        // Scale the data.
	InverseTransform(v float64) float64 // Unscale the data.
	GetParam(key string) *float64       // Returns the parameter for the given key, and nil if the parameter does not exist.
}

// #region MinMaxScaler
type MinMaxScaler struct {
	min float64 // Minimum value
	max float64 // Maximum value
}

// NewMinMaxScaler returns a new instance of the MinMaxScaler.
func NewMinMaxScaler() *MinMaxScaler {
	return &MinMaxScaler{}
}

func (s *MinMaxScaler) Fit(data []float64) {
	s.min = floats.Min(data)
	s.max = floats.Max(data)
}

func (s *MinMaxScaler) Transform(v float64) float64 {
	return (v - s.min) / (s.max - s.min)
}

func (s *MinMaxScaler) InverseTransform(v float64) float64 {
	return v*(s.max-s.min) + s.min
}

func (s *MinMaxScaler) GetParam(key string) *float64 {
	switch key {
	case "max":
		return &s.max
	case "min":
		return &s.min
	default:
		return nil
	}
}

//#endregion

// #region StandardScaler
type StandardScaler struct {
	mean              float64
	standardDeviation float64
}

// StandardScaler returns a new instance of the StandardScaler.
func NewStandardScaler() *StandardScaler {
	return &StandardScaler{}
}

func (s *StandardScaler) Fit(data []float64) {
	mu, _ := stats.Mean(data)
	sd, _ := stats.StandardDeviation(data)
	s.mean = mu
	s.standardDeviation = sd
}

func (s *StandardScaler) Transform(v float64) float64 {
	return (v - s.mean) / s.standardDeviation
}

func (s *StandardScaler) InverseTransform(v float64) float64 {
	return v*s.standardDeviation + s.mean
}

func (s *StandardScaler) GetParam(key string) *float64 {
	switch key {
	case "mean":
		return &s.mean
	case "stdev":
		return &s.standardDeviation
	default:
		return nil
	}
}

//#endregion
