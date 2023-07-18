package main

import "fmt"

// Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
//Print the decimal value of each fraction on a new line with  places after the decimal.

func plusMinus(arr []int32) {
	size := len(arr)
	positives := []int32{}
	negatives := []int32{}
	zeros := []int32{}

	for _, number := range arr {
		if number == 0 {
			zeros = append(zeros, number)
		} else if number > 0 {
			positives = append(positives, number)
		} else if number < 0 {
			negatives = append(negatives, number)
		}
	}

	fmt.Println(float64(len(positives)) / float64(size))
	fmt.Println(float64(len(negatives)) / float64(size))
	fmt.Println(float64(len(zeros)) / float64(size))
}

func main() {
	testNumbers := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(testNumbers)

	// Expected Output
	// 0.500000
	// 0.333333
	// 0.166667
}
