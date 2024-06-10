package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	count := 0
	for _, char := range input {
		if char != 'a' && char != 'A' && char != 'i' && char != 'I' && char != '.' {
			count++
		} else if char == '.' {
			break
		}
	}

	fmt.Println(count)
}