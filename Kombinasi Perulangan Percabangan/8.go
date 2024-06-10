package main

import (
	"fmt"
)

func isPrimeX(n int) bool {
	factors := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors++
		}
		if factors > 3 {
			return false
		}
	}
	return factors == 3
}

func main() {
	var num int

	fmt.Scan(&num)
	fmt.Println(isPrimeX(num))
}