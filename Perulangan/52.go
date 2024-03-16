package main

import "fmt"

func main() {
	var numbers [10]int
	var input, lastNumber, count int
	fmt.Scan(&lastNumber)
	numbers[0] = lastNumber
	count = 1
	for i := 1; i < 10; i++ {
		fmt.Scan(&input)

		if input > lastNumber {
			break
		}

		numbers[i] = input
		lastNumber = input
		count++
	}

	fmt.Printf("%d\n", count)
}
