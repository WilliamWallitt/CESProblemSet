package main

import "fmt"

func main () {
	fmt.Println(CoinPiles(2, 1), TestPiles(2, 1))
	fmt.Println(CoinPiles(2, 2), TestPiles(2, 2))
	fmt.Println(CoinPiles(3, 3), TestPiles(3, 3))
	fmt.Println(CoinPiles(500, 1000), TestPiles(500, 1000))

}


func CoinPiles (a int, b int) string {

	counter := 0
	for a > 0 && b > 0 {

		if a > b {
			a -= 2
			b -= 1
		} else if b < a {
			a -= 1
			b -= 2
		} else {
			a -= 1
			b -= 2
		}

		counter ++

	}

	if a != 0 || b != 0 {
		return "NO"
	} else {
		return "YES"

	}

}


func TestPiles (a int, b int) string {

	if a < b {
		inter := a
		a = b
		b = inter
	}

	if a > 2*b {
		return "NO"
	}

	if (a + b)%3 == 0 {
		return "YES"
	} else {
		return "NO"
	}


}