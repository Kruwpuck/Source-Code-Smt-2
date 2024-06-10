package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanln(&input)

	sum := 0
	for _, char := range input {
		if char >= '0' && char <= '9' {
			num := int(char - '0')
			sum += num
		}
	}

	fmt.Println(sum)
}