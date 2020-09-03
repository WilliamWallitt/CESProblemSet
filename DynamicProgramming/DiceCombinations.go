package main

import (
	"fmt"
)

func main() {
	test := diceCombinations(6)
	fmt.Println(test)
}


func diceCombinations(n int) int {

	poss, index, counter := make([]int, n), 0, 0


	for index < len(poss) {

		if poss[index] > 5 {
			poss[index] = 1
			index ++
		}

		if index > len(poss) - 1 {
			break
		}

		poss[index] = poss[index] + 1

		currSum := poss[index]
		for i := 0; i < len(poss); i++ {
			if index != i {
				currSum += poss[i]
			}
		}

		if currSum == n {
			fmt.Println(poss)
			counter ++
		}

	}

	return counter

}