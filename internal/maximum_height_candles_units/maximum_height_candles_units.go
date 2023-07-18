package maximum_height_candles_units

// You are in charge of the cake for a child's birthday.
// You have decided the cake will have one candle for each year of their total age.
// They will only be able to blow out the tallest of the candles. Count how many candles are tallest.
func birthdayCakeCandles(candles []int32) int32 {
	var tallestCandleHight int32 = 0

	for _, candle := range candles {
		if candle > tallestCandleHight {
			tallestCandleHight = candle
		}
	}

	var tallestCandleCounter int32 = 0

	for _, candle := range candles {
		if candle == tallestCandleHight {
			tallestCandleCounter += 1
		}
	}

	return tallestCandleCounter
}
