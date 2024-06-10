package main

import (
	"fmt"
)

func main() {
	var num, sum int

	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		if num%3 == 0 && num%2 != 0 {
			sum += num
		}
	}

	fmt.Printf("%d\n", sum)
}