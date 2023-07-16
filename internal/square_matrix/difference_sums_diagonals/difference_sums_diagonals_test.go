package difference_sums_diagonals

import "testing"

func TestDiagonalDifference(t *testing.T) {
	arr1 := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	expected1 := int32(2)
	result1 := diagonalDifference(arr1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	arr2 := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	expected2 := int32(15)
	result2 := diagonalDifference(arr2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	arr3 := [][]int32{
		{10, 20},
		{30, 40},
	}
	expected3 := int32(0)
	result3 := diagonalDifference(arr3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}
}
