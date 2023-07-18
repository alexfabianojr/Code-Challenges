package main

import "fmt"

// Print a staircase as described
func staircase(n int32) {
	var stepCounter int32 = 1
	for {
		var spaceCounter int32 = 1

		for {
			if spaceCounter > n {
				break
			} else if spaceCounter > n-stepCounter {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
			spaceCounter += 1
		}

		if stepCounter == n {
			break
		}

		fmt.Println("")
		stepCounter += 1
	}
}

func main() {
	staircase(6)
	// Expected
	//	   #
	//    ##
	//   ###
	//  ####
	// #####
	//######
}
