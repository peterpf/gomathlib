package num

import "testing"

func TestTotient(t *testing.T) {
	var values = []int {7, 8}
	var expectedResults = [][]int {{1, 2, 3, 4, 5, 6}, {1, 3, 5, 7}}

	for i := 0; i<len(values); i++ {
		result := Totient(values[i])
		expectedResult := expectedResults[i]
		if len(expectedResult) != len(result) {
			t.Errorf("Expected %v but got %v", expectedResult, result)
			return
		}
		for j := 0; j<len(result); j++ {
			if expectedResult[j] != result[j] {
				t.Errorf("Expected %d but got %d of input %v at index %d", expectedResult[j], result[j], expectedResult, j)
			}
		}
	}
}
