package num

import (
	"testing"
)

func TestTotient(t *testing.T) {
	var inputValues = []int {7, 8}
	var expectedResults = [][]int {{1, 2, 3, 4, 5, 6}, {1, 3, 5, 7}}

	for i := 0; i<len(inputValues); i++ {
		result := Totient(inputValues[i])
		expectedResult := expectedResults[i]
		if len(expectedResult) != len(result) {
			t.Errorf("Expected %v but got %v", expectedResult, result)
			return
		}
		for j := 0; j<len(result); j++ {
			if expectedResult[j] != result[j] {
				t.Errorf("Expected %d but got %d at index %d for Totient(%d)", expectedResult[j], result[j], j, expectedResult)
			}
		}
	}
}

func TestMoebius(t *testing.T) {
	var inputValues = []int {4, 7, 15}
	var expectedResults = []int {0, -1, 1}
	for i := 0; i < len(inputValues); i++ {
		result := Moebius(inputValues[i])
		expectedResult := expectedResults[i]
		if expectedResult != result {
			t.Errorf("Expected %d but got %d for Moebius(%d)", expectedResult, result, inputValues[i])
		}
	}
}

func TestGCD(t *testing.T) {
	var inputValues = [][]int {{1, 7}, {4, 8}, {7, 13}, {49865, 69811}}
	var expectedResults = []int {1, 4, 1, 9973}
	for i := 0; i < len(inputValues); i++ {
		a := inputValues[i][0]
		b := inputValues[i][1]
		result := GCD(a, b)
		expectedResult := expectedResults[i]
		if expectedResult != result {
			t.Errorf("Expected %d but got %d for GCD(%d, %d)", expectedResult, result, a, b)
		}
	}
}

func TestPF(t *testing.T) {
	var inputValues = []int {1, 4, 7, 10, 15}
	var expectedResults = [][]int {{}, {2, 2}, {7}, {2, 5}, {3, 5}}
	for i := 0; i < len(inputValues); i++ {
		result := PF(inputValues[i])
		expectedResult := expectedResults[i]
		if len(expectedResult) != len(result) {
			t.Errorf("Expected %v but got %v for PF(%d)", expectedResult, result, inputValues[i])
			return
		}
		for j :=0; j < len(expectedResult); j++ {
			if expectedResult[j] != result[j] {
				t.Errorf("Expected %v but got %v for PF(%d)", expectedResult, result, inputValues[j])
			}
		}
	}
}

func TestPFExponent(t *testing.T) {
	var inputValues = [][]int {{4, 2}, {8, 2}, {15, 3}}
	var expectedResults = []int {2, 3, 1}
	for i := 0; i < len(inputValues); i++ {
		n := inputValues[i][0]
		p := inputValues[i][1]
		result := PFExponent(n, p)
		expectedResult := expectedResults[i]
		if expectedResult != result {
			t.Errorf("Expected %d but got %d for PFExponent(%d, %d)", expectedResult, result, n, p)
		}
	}
}
