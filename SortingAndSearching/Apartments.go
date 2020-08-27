package main

import "fmt"

func main() {

	n, m, k, desiredSize, actualSize := 6, 5, 5, []int{60, 55, 80, 60, 100, 90}, []int{85, 60, 75, 50, 10}

	fmt.Println(Apartments(n, m, k, desiredSize, actualSize))

}


func Apartments (n int, m int, k int, desiredSize []int, actualSize []int) int {

	taken := make([]int, actualSize[0])
	counter := 0

	for i := 0; i < m; i++ {

		currApartment, lowerBound, upperBound := actualSize[i], actualSize[i] - k, actualSize[i] + k

		for j := 0; j < n; j++ {

			currClient := desiredSize[j]

			if currClient >= lowerBound && currClient <= upperBound {

				for k := 0; k < len(taken); k++ {

					if currApartment == taken[k] {
						break
					}

					if k == len(taken) - 1 {
						counter += 1
						taken = append(taken, currApartment)
					}

				}
			}
		}

	}

	return counter

}