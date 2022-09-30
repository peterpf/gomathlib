package gomathlib

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestAverageFilter(t *testing.T) {
	// Arrange
	filter := NewAverageFilter(1)
	dataMatrix := NewMatrixWithValue(4, 4, 1)
	expectedValues := []float64{1 / 9.0, 2 / 9.0, 2 / 9.0, 2 / 9.0, 2 / 9.0, 4 / 9.0, 4 / 9.0, 4 / 9.0, 2 / 9.0, 4 / 9.0, 4 / 9.0, 4 / 9.0, 2 / 9.0, 4 / 9.0, 4 / 9.0, 4 / 9.0}
	expectedMatrix := mat.NewDense(4, 4, expectedValues)

	// Act
	resultMatrix := filter.Apply(dataMatrix)

	// Assert
	fmt.Print(resultMatrix.RawMatrix().Data)
	if expectedMatrix.RawMatrix().Rows != resultMatrix.RawMatrix().Rows {
		t.Errorf("expected %d rows, got %d", expectedMatrix.RawMatrix().Rows, resultMatrix.RawMatrix().Rows)
	}
	if expectedMatrix.RawMatrix().Cols != resultMatrix.RawMatrix().Cols {
		t.Errorf("expected %d columns, got %d", expectedMatrix.RawMatrix().Cols, resultMatrix.RawMatrix().Cols)
	}
	resultMatrix.Apply(func(i, j int, v float64) float64 {
		if expectedMatrix.At(i, j) != v {
			t.Errorf("expected value %f at (%d, %d), got %f", expectedMatrix.At(i, j), i, j, v)
		}
		return v
	}, resultMatrix)
}

func TestGaussFilter(t *testing.T) {
	// Arrange
	sigma := 1.0
	filter := NewGaussianFilter(1, &sigma)
	dataMatrix := NewMatrixWithValue(4, 4, 1)
	expectedValues := []float64{0.07511360795411151, 0.1989550111070855, 0.1989550111070855, 0.1989550111070855, 0.1989550111070855, 0.5269763698317176, 0.5269763698317176, 0.5269763698317176, 0.1989550111070855, 0.5269763698317176, 0.5269763698317176, 0.5269763698317176, 0.1989550111070855, 0.5269763698317176, 0.5269763698317176, 0.5269763698317176}
	expectedMatrix := mat.NewDense(4, 4, expectedValues)

	// Act
	resultMatrix := filter.Apply(dataMatrix)

	// Assert
	if expectedMatrix.RawMatrix().Rows != resultMatrix.RawMatrix().Rows {
		t.Errorf("expected %d rows, got %d", expectedMatrix.RawMatrix().Rows, resultMatrix.RawMatrix().Rows)
	}
	if expectedMatrix.RawMatrix().Cols != resultMatrix.RawMatrix().Cols {
		t.Errorf("expected %d columns, got %d", expectedMatrix.RawMatrix().Cols, resultMatrix.RawMatrix().Cols)
	}
	resultMatrix.Apply(func(i, j int, v float64) float64 {
		if expectedMatrix.At(i, j) != v {
			t.Errorf("expected value %f at (%d, %d), got %f", expectedMatrix.At(i, j), i, j, v)
		}
		return v
	}, resultMatrix)
}
