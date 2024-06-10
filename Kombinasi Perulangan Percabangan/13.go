package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	countA := 0
	countI := 0

	for i := 0; i < len(input); i++ {
		if input[i] == 'a' || input[i] == 'A' {
			countA++
		} else if input[i] == 'i' || input[i] == 'I' {
			countI++
		} else if input[i] == '.' {
			break
		}
	}

	totalCount := countA + countI
	fmt.Println(totalCount)
}