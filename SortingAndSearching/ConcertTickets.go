package main

import (
	"fmt"
	"sort"
)

func main() {

	test(5, 3, []int{5, 3, 7, 8, 5}, []int{4, 8, 3})

}


func concertTickets(n int, m int, priceOfTicket []int, customerPrice []int) {


	sort.Ints(priceOfTicket)
	taken := append(make([]int, 0), -1)

	for i := 0; i < len(customerPrice); i++ {
		currCustomer := customerPrice[i]
		found := false
		for j := len(priceOfTicket) - 1; j >= 0; j-- {

			if found {break}

			if priceOfTicket[j] <= currCustomer {

				for k := 0; k < len(taken); k++ {
 					currItem := taken[k]
					if currItem == priceOfTicket[j] {
						break
					}
					if k == len(taken) - 1 {
						taken = append(taken, priceOfTicket[j])
						fmt.Println("Ticket: ", priceOfTicket[j], "Customer: ", currCustomer)
						found = true
					}
				}
			}

			if j == 0 && !found {
				fmt.Println("Ticket: ", -1, "Customer: ", currCustomer)
				break
			}

		}
	}

}

func test(n int, m int, priceOfTicket []int, customerPrice []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(priceOfTicket)))
	for _, num := range customerPrice {
		index := 0
		for _, ticket := range priceOfTicket{
			if ticket <= num {
				fmt.Println(ticket)
				priceOfTicket = remove(priceOfTicket, index)
				break
			}
			if index == len(priceOfTicket) - 1 {
				fmt.Println(-1)
				break
			}
			index ++
		}
	}
}

func remove(s []int, i int) []int {

	s[len(s) - 1], s[i] = s[i], s[len(s) - 1]
	return s[:len(s) - 1]
}
