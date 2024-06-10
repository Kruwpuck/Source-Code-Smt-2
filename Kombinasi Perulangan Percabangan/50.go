package main

import (
	"fmt"
)

func main() {
	var num, sum int
	// Loop while-like using for in Go
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		if num%4 == 0 && num%2 == 0 {
			sum += num
		}
	}

	fmt.Print(sum)
}