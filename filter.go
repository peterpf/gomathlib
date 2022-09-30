package gomathlib

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Filter interface {
	Apply(data *mat.Dense) *mat.Dense
}

// #region Average Filter
type AverageFilter struct {
	kernel *mat.Dense
}

// NewAverageFilter constructs a kernel of size 2*k+1, resulting in a filter matrix of size (2k+1)x(2k+1).
func NewAverageFilter(k int) *AverageFilter {
	filterSize := 2*k + 1
	cellValue := 1 / float64(filterSize*filterSize)
	return &AverageFilter{
		kernel: NewMatrixWithValue(filterSize, filterSize, cellValue),
	}
}

func (f *AverageFilter) Apply(data *mat.Dense) *mat.Dense {
	return apply(data, f.kernel)
}

//#endregion

// #region GaussianFilter
type GaussianFilter struct {
	kernel *mat.Dense
}

// NewGaussianFilter constructs a gauss kernel of size 2*k+1, resulting in a filter matrix of size (2k+1)x(2k+1).
func NewGaussianFilter(k int, sigma *float64) *GaussianFilter {
	filterSize := 2*k + 1
	radius := filterSize / 2
	sd := math.Max(float64(radius)/2, 1)
	if sigma != nil {
		sd = *sigma
	}
	kernel := mat.NewDense(filterSize, filterSize, nil)

	// Populate filter with values
	kernel.Apply(func(i, j int, v float64) float64 {
		return gaussian(float64(i-radius), float64(j-radius), sd)
	}, kernel)

	// Normalize kernel so that the kernel values add up to 1.
	sum := mat.Sum(kernel)
	kernel.Apply(func(i, j int, v float64) float64 { return v / sum }, kernel)

	return &GaussianFilter{
		kernel: kernel,
	}
}

// gaussian returns the value for given x and y coordinates depending on sigma.
func gaussian(x, y, sigma float64) float64 {
	return 1 / (2 * math.Pi * math.Pow(sigma, 2)) * math.Exp(-(math.Pow(x, 2)+math.Pow(y, 2))/(2*math.Pow(sigma, 2)))
}

func (f *GaussianFilter) Apply(data *mat.Dense) *mat.Dense {
	return apply(data, f.kernel)
}

//#endregion

//#region Helper Functions

// apply runs the kernel over the input data matrix and returns the result as a new matrix.
func apply(data *mat.Dense, kernel *mat.Dense) *mat.Dense {
	result := mat.DenseCopyOf(data)
	nrows := data.RawMatrix().Rows
	ncols := data.RawMatrix().Cols
	radius := kernel.RawMatrix().Rows / 2
	for rowIdx := 0; rowIdx < nrows; rowIdx++ {
		for colIdx := 0; colIdx < ncols; colIdx++ {
			lowerboundRowIdx := int(math.Max(float64(rowIdx-radius), 0))
			upperboundRowIdx := int(math.Min(float64(rowIdx+radius), float64(nrows)))
			lowerboundColIdx := int(math.Max(float64(colIdx-radius), 0))
			upperboundColIdx := int(math.Min(float64(colIdx+radius), float64(ncols)))
			var sum float64
			for i := lowerboundRowIdx; i < upperboundRowIdx; i++ {
				for j := lowerboundColIdx; j < upperboundColIdx; j++ {
					filterValue := kernel.At(i-lowerboundRowIdx, j-lowerboundColIdx)
					sum += filterValue * data.At(i, j)
				}
			}
			result.Set(rowIdx, colIdx, sum)
		}
	}
	return result
}

// NewMatrixWithValue returns a new matrix populated with a specific value.
func NewMatrixWithValue(nrows, ncols int, fillValue float64) *mat.Dense {
	data := make([]float64, nrows*ncols)
	for i := 0; i < len(data); i++ {
		data[i] = fillValue
	}
	return mat.NewDense(nrows, ncols, data)
}

//#endregion
