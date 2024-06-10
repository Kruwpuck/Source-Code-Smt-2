package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i:i+2] == "go" {
			count++
		}
	}

	fmt.Println(count)
}