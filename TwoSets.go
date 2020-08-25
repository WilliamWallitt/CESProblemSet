package main

import "fmt"

func main () {

	TwoSets(7)

}

func TwoSets (number int) {

	var sum int = 0

	for i := 1; i <= number; i++ {
		sum += i
	}

	if sum % 2 == 0 {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")
	}



}
