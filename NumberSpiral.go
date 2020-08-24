package main

import "fmt"

func main () {

	NumberSpiral(3, 5, 5)

}


func NumberSpiral (tests int, y int, x int) int {

	var startingNum = 1
	var layerAmount int

	if y < x {
		layerAmount = x
	} else {
		layerAmount = y
	}

	var coordinates [][]int


	for i := 1; i <= layerAmount; i++ {

		if len(coordinates) > 0 {
			break
		}

		var layer = getLayerCoords(i)

		startingNum += len(layer)

		if layer[0][0] == layerAmount || layer[0][1] == layerAmount {
			coordinates = layer
			startingNum -= len(layer)
		}
	}

	for i := 0; i < len(coordinates); i++ {

		if layerAmount % 2 == 0 {
			if coordinates[i][0] == x && coordinates[i][1] == y {
				fmt.Println(startingNum)
				break
			}
		} else {
			if coordinates[i][0] == y && coordinates[i][1] == x {
				fmt.Println(startingNum)
				break
			}
		}
		startingNum += 1
	}

	return 0


}

func getLayerCoords(layer int) [][]int {

	var coords = make([][]int, layer*2 - 1)
	x, y := layer, 1
	index := 0
	for x > 0 {
		coords[index] = []int{x, y}
		if y == layer {
			x --
		} else {
			y ++
		}
		index ++
	}
	return coords

}


// ****
// working

//for j := 0; j < len(layer); j++ {
//
//	if layer[j][0] == x && layer[j][1] == y {
//		coordinates = layer
//		startingNum -= len(layer)
//	}
//
//}
