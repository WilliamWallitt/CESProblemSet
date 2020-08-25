package main

import (
	"fmt"
	"strconv"
)

func main () {

	// need to implement array to store value Uint64 wont work
	res := TrailingZeros(20)
	fmt.Println(res)
}


func TrailingZeros (n int) int {

	num := strconv.FormatUint(Factorial(n), 10)
	counter := 0
	for i := len(num) - 1; i >= 0; i-- {
		if string(num[i]) == "0" {
			counter ++
		} else {
			break
		}
	}

	return counter

}

func Factorial(n int) uint64 {
	res := uint64(n)
	if n == 1 {
		return res
	} else {
		res *= Factorial(n - 1)
	}
	return res

}