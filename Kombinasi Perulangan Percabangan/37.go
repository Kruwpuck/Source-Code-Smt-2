package main

import "fmt"

func main() {
	var totalEvenTriplets int

	for {
		var dice1, dice2, dice3 int
		fmt.Scanln(&dice1, &dice2, &dice3)

		if dice1+dice2+dice3 == 12 {
			break
		}

		if dice1%2 == 0 && dice2%2 == 0 && dice3%2 == 0 {
			totalEvenTriplets++
		}
	}

	fmt.Println(totalEvenTriplets)
}