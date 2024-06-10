package main

import (
	"fmt"
)

func main() {
	var num int
	var countPositive, countNegative int

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			countPositive++
		} else if num < 0 {
			countNegative++
		}
	}

	fmt.Print(countPositive, " ")
	fmt.Print(countNegative, " ")
}