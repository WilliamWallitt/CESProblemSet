package main

import "fmt"


var global = make([]string, 0)

func main() {

	set := []rune("aabac")
	CreatingStrings1(set, 0, len(set) - 1)

	check := make(map[string]int)
	res := make([]string, 0)

	// create [string, 1] object
	// so only unique values are in check
	for _, val := range global {
		check[val] = 1
	}
	// populate result array
	for letter, _ := range check {
		res = append(res, letter)
	}
	fmt.Println(res, len(res))

}

func CreatingStrings1 (set []rune, index int, right int) {

	if index > len(set) {

		global = append(global, string(set))
		return

	} else {

		CreatingStrings1(set, index + 1, right)

		for i := index + 1; i < len(set); i++ {

			set = swap(set, index, i)
			CreatingStrings1(set, index + 1, right)
			set = swap(set, index, i)

		}
	}

}

func swap (s []rune, i int, j int) []rune {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
	return s

}
