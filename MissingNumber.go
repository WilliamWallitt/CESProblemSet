package main

import (
	"fmt"
)

func main() {

	x := missingNumber(20, []int{1,2,3,4,5,6,7,8,10,11,12,13,14,15,16,17,18, 19,20})
	// 9 is missing
	fmt.Println(x)
}


func missingNumber (n int, arr []int) int {

	var missing int

	for i := 1; i <= n; i++ {
		for j := 0; j < len(arr); j++ {
			curr := arr[j]

			if curr == i {
				break
			}

			if j == len(arr) - 1 {
				missing = i
				return missing
			}

		}
	}

	return missing

}
