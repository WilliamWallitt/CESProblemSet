package main

import "fmt"

func main() {
	Permutations(10)
}

func Permutations(n int) {

	slice := make([]int, n)

	for i := n; i > 0; i-- {

		num := i

		for j := 0; j < n; j++ {

			item := slice[j]

			if j == 0 {

				next := slice[j + 1]

				if item == 0 && num != next && num - 1 != next && num + 1 != next {
					slice[j] = num
					break
				}


			} else if j > 0 && j < n - 1 {

				prev := slice[j - 1]
				next := slice[j + 1]

				if item == 0 && num != next && num != prev && num - 1 != prev && num - 1 != next && num + 1 != prev &&
					num + 1 != next {
					slice[j] = num
					break
				}


			} else {

				prev := slice[j - 1]

				if item == 0 && num != prev && num + 1 != prev && num -1 != prev {
					slice[j] = num
					break

				}


			}

		}
	}

	for i := 0; i < len(slice); i++ {

		if slice[i] == 0 {
			fmt.Println("No Solutions")
			break
		}

		if i == len(slice) - 1 && slice[i] != 0 {
			fmt.Println(slice)
		}
	}

}