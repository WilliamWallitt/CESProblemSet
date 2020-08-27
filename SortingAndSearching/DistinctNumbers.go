package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := []int{2,3,2,2,3,4,4,4,5,5,2}
	fmt.Println(distinctNumbers(numbers))

}


func distinctNumbers(nums []int) int {

	sort.Ints(nums)
	curr, index, counter := nums[0], 0, 1
	for i := 0; i < len(nums); i++ {
		if i != index && nums[i] != curr {
			counter ++
			curr = nums[i]
			index = i
		}
	}

	return counter
}
