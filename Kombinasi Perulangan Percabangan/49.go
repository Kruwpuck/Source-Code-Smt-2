package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)

	fmt.Scan(&m)

	result := 0
	for i := 0; i < n; i++ {
		result += m
	}

	fmt.Print(result)
}