package maximum_height_candles_units

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	candles1 := []int32{4, 2, 4, 4, 3}
	expected1 := int32(3)
	result1 := birthdayCakeCandles(candles1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	candles2 := []int32{1, 1, 1, 1}
	expected2 := int32(4)
	result2 := birthdayCakeCandles(candles2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	candles3 := []int32{5, 4, 3, 2, 1}
	expected3 := int32(1)
	result3 := birthdayCakeCandles(candles3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}
}
