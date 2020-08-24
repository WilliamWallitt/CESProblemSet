package main

import "fmt"

func main() {
	x := IncreasingArray(5, []int{3,2,5,1,7})
	fmt.Println(x)
}

func IncreasingArray(size int, arr []int) int {

	var moves int = 0

	for i := 0; i < size -1; i++ {
		curr := arr[i]
		if arr[i + 1] < curr {
			moves += curr - arr[i + 1]
			arr[i + 1] = curr
		}

	}

	return moves

}
