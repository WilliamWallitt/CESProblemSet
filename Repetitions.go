package main

import "fmt"

func main() {

	x := Repetitions("ATTCCCCCCGGGGGAAAATTTTTTCCCGGGTTTTGGGGAAAAAAAAAAAAAAA")
	fmt.Println(x)
}

func Repetitions(sequence string) int {

	currentChar := string(sequence[0])
	longest := 0
	currRep := 1

	for i := 1; i < len(sequence); i++ {

		// start at index 1

		current := string(sequence[i])

		if current == currentChar {

			currRep = currRep + 1

			if i == len(sequence) - 1 {
				if currRep > longest {
					longest = currRep
				}
			}

		} else {

			if currRep > longest {
				longest = currRep
			}
			currRep = 1
			currentChar = string(sequence[i])
		}

	}

	return longest

}