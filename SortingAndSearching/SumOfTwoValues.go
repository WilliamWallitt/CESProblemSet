package main

import "fmt"

func main() {

	sumOfTwoValues(6, []int{2,7,5,1})

}

func sumOfTwoValues(sum int, numbers []int) {

	// finding index of integers that add up to the target (sum)
	// so [2,7,5,1] with target 8 would be index [2, 4] or numbers 5 and 1

	currIndex, index, answers := 0, 1, make([][]int, 0)

	for true {
		if index > len(numbers) - 1  {
			currIndex ++
			index = 0
		}
		if currIndex > len(numbers) - 1 {
			break
		}
		if index != currIndex {
			if numbers[index] + numbers[currIndex] == sum {
				answers = append(answers, []int{currIndex, index})
			}
		}
		index ++
	}

	for _, solution := range answers {
		fmt.Println(solution[0] + 1, solution[1] + 1)
	}

}