package main

import "fmt"

func main() {
	boards(100)
}

func boards (n int) {
	for i := 0; i < n; i++ {
		fmt.Println(twoKnights(i + 1))
	}
}

func twoKnights(n int) int {

	var coordinates = generateCoords(n)
	var k1 []int
	var k2 []int
	var pos int

	for i := 0; i < len(coordinates); i++ {
		k1 = coordinates[i]
		for j := 0; j < len(coordinates); j++ {
			if j != i {
				k2 = coordinates[j]
				x1, y1, x2, y2 := k1[0], k1[1], k2[0], k2[1]
				// possible moves
				if x2 + 2 == x1 && y2 + 1 == y1 || x2 - 2 == x1 && y2 + 1 == y1 || x2 + 1 == x1 && y2 + 2 == y1 ||
					x2 - 1 == x1 && y2 + 2 == y1 || x2 + 2 == x1 && y2 - 1 == y1 || x2 - 2 == x1 && y2 - 1 == y1 ||
					x2 + 1 == x1 && y2 - 2 == y1 || x2 - 1 == x1 && y2 - 2 == y1 {
					continue
				} else {
					pos ++
				}
			}
		}
	}
	return pos / 2
}

func generateCoords(n int) [][]int {
	var coords = make([][]int, n*n)
	i, x, y := 0, 1, 1
	for i < n*n {
		coords[i] = []int{x, y}
		if x >= n {
			y ++
			x = 1
		} else {
			x++
		}
		i ++
	}
	return coords
}