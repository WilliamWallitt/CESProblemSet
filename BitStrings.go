package main

import (
	"fmt"
	"math"
)

func main() {

	res := BitStrings(150)
	fmt.Println(res)
}

func BitStrings (n int) uint64 {

	return uint64(math.Pow(float64(2), float64(n)))

}
