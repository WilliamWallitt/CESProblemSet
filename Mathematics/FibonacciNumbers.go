package main

import "fmt"

func main () {

	test := spaceFibonacci(150)
	fmt.Println(test)

}

// simple recursive function
// complexity is T(n) = T(n - 1) + T( n - 2) ... therefore exponential
func fibonacci (n int) int {

	if n == 0 {
		return 0
	} else {

		if n == 1 {
			return 1
		} else {

			sum := fibonacci(n - 1) + fibonacci(n - 2)
			return sum

		}

	}
}

func spaceFibonacci (n int) int {

	a, b := 0, 1
	if n < 0 {
		return a
	} else if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else {
		// using a sliding window - only store prev two values
		for i := 2; i < n + 1; i++ {
			// create next fibonacci number
			c := a + b
			// a now is the next term (b)
			a = b
			// b is now the current term (c)
			b = c
			// repeat
		}
		// return the last term calculated
		return b
	}

}
