package gomathlib

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestMinMaxScaler(t *testing.T) {
	// Arrange
	scaler := NewMinMaxScaler()
	data := []float64{1, 2, 3, 4}
	min := 1.0
	max := 4.0
	expectedData := []float64{0, 1 / 3.0, 2 / 3.0, 3 / 3.0}

	//Act
	scaler.Fit(data)

	// Assert
	if scaler.GetParam("min") == nil && *scaler.GetParam("min") != min {
		t.Errorf("expected `min` parameter to be %f, got %v", min, scaler.GetParam("min"))
	}
	if scaler.GetParam("max") == nil && *scaler.GetParam("max") != max {
		t.Errorf("expected `max` parameter to be %f, got %v", max, scaler.GetParam("max"))
	}
	if scaler.GetParam("unknown_param") != nil {
		t.Errorf("expected `unknown_param` parameter to be nil, got %v", scaler.GetParam("unknown_param"))
	}

	for i := 0; i < len(data); i++ {
		// Check transform
		scaledValue := scaler.Transform(data[i])
		if expectedData[i] != scaledValue {
			t.Errorf("expected %f to be scaled to %f, got %f", data[i], expectedData[i], scaledValue)
		}
		// Check inverseTransform
		unscaledValue := scaler.InverseTransform(scaledValue)
		if unscaledValue != data[i] {
			t.Errorf("expected the inverse of %f to be %f, got %f", scaledValue, data[i], unscaledValue)
		}
	}
}

func TestStandardScaler(t *testing.T) {
	// Arrange
	scaler := NewStandardScaler()
	data := []float64{1, 2, 3, 4}
	mean, _ := stats.Mean(data)
	stdev, _ := stats.StandardDeviation(data)
	expectedData := []float64{(1 - mean) / stdev, (2 - mean) / stdev, (3 - mean) / stdev, (4 - mean) / stdev}

	//Act
	scaler.Fit(data)

	// Assert
	if scaler.GetParam("mean") == nil && *scaler.GetParam("mean") != mean {
		t.Errorf("expected `mean` parameter to be %f, got %v", mean, scaler.GetParam("mean"))
	}
	if scaler.GetParam("stdev") == nil && *scaler.GetParam("stdev") != stdev {
		t.Errorf("expected `stdev` parameter to be %f, got %v", stdev, scaler.GetParam("stdev"))
	}
	if scaler.GetParam("unknown_param") != nil {
		t.Errorf("expected `unknown_param` parameter to be nil, got %v", scaler.GetParam("unknown_param"))
	}

	for i := 0; i < len(data); i++ {
		// Check transform
		scaledValue := scaler.Transform(data[i])
		if expectedData[i] != scaledValue {
			t.Errorf("expected %f to be scaled to %f, got %f", data[i], expectedData[i], scaledValue)
		}

		// Check inverseTransform
		unscaledValue := scaler.InverseTransform(scaledValue)
		// Round the values due to floating point inaccuracies when unscaling data, and thus failing tests.
		if math.Round(unscaledValue*10e6)/10e6 != math.Round(data[i]*10e6)/10e6 {
			t.Errorf("expected the inverse of %f to be %f, got %f", scaledValue, data[i], unscaledValue)
		}
	}
}

func TestLimiter(t *testing.T) {
	// Arrange
	min := 0.0
	max := 255.0
	scaler := NewLimiter(min, max)
	data := []float64{-4, 255, 256, 42}
	expectedData := []float64{min, 255, max, 42}

	//Act
	scaler.Fit(data)

	// Assert
	if scaler.GetParam("min") == nil && *scaler.GetParam("min") != min {
		t.Errorf("expected `min` parameter to be %f, got %v", min, scaler.GetParam("min"))
	}
	if scaler.GetParam("max") == nil && *scaler.GetParam("max") != max {
		t.Errorf("expected `max` parameter to be %f, got %v", max, scaler.GetParam("max"))
	}
	if scaler.GetParam("unknown_param") != nil {
		t.Errorf("expected `unknown_param` parameter to be nil, got %v", scaler.GetParam("unknown_param"))
	}

	for i := 0; i < len(data); i++ {
		// Check transform
		scaledValue := scaler.Transform(data[i])
		if expectedData[i] != scaledValue {
			t.Errorf("expected %f to be scaled to %f, got %f", data[i], expectedData[i], scaledValue)
		}

		// Check inverseTransform
		unscaledValue := scaler.InverseTransform(scaledValue)
		// Round the values due to floating point inaccuracies when unscaling data, and thus failing tests.
		if unscaledValue != scaledValue {
			t.Errorf("expected the inverse of %f to be %f, got %f", scaledValue, data[i], unscaledValue)
		}
	}
}
