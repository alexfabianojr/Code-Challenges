package difference_sums_diagonals

// Given a square matrix, calculate the absolute difference between the sums of its diagonals.
func diagonalDifference(arr [][]int32) int32 {
	diagonals := []int32{}

	for line, lineValue := range arr {
		diagonals = append(diagonals, lineValue[line])
	}

	var leftDiagonal int32 = 0

	for _, diagonal := range diagonals {
		leftDiagonal += diagonal
	}

	diagonals = []int32{}

	for line, lineValue := range arr {
		size := len(lineValue) - 1
		if size-line >= 0 {
			diagonals = append(diagonals, lineValue[size-line])
		}
	}

	var rightDiagonal int32 = 0

	for _, diagonal := range diagonals {
		rightDiagonal += diagonal
	}

	difference := leftDiagonal - rightDiagonal

	if difference < 0 {
		return difference * -1
	}

	return difference
}
