package main

import "fmt"

func main() {

	test := restaurantCustomers([][]int{{5,8}, {2,4}, {3,9}, {1, 11}})
	fmt.Println(test)
}


func restaurantCustomers(customerTimes [][]int) int {
	maxCustomer := 0
	for index, currTime := range customerTimes{
		startTime, endTime, counter := currTime[0], currTime[1], 0
		for i, nextTime := range customerTimes {
			if index != i {
				st, et := nextTime[0], nextTime[1]
				if st >= startTime || endTime <= et {
					counter ++
				}
			}
		}
		if counter > maxCustomer {
			maxCustomer = counter
		}

	}

	return maxCustomer


}