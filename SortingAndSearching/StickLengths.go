package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	test := sticks([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(test)

}

func sticks (stickLengths []int) int {
	if len(stickLengths) < 2 {
		return stickLengths[0]
	}
	sort.Ints(stickLengths)
	// [min, maxx]
	median, counter := 0, 0
	for _, value := range stickLengths {
		median += value
	}
	median = median / len(stickLengths)
	for _, item := range stickLengths {
		if item < median {
			counter += item - median
		} else {
			counter += median - item
		}
	}
	return int(math.Abs(float64(counter)))
}