package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x string = wierdAlgorithm(23)
	fmt.Println(x)
}



func wierdAlgorithm(n int) string {
	var seq string

	for n != 1 {
		seq += strconv.Itoa(n) + " -> "
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = (n*3) + 1
		}
	}

	seq += strconv.Itoa(1)

	return seq

}