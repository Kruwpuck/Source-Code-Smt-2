package main

import "fmt"

func isConsonant(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', '.':
		return false
	default:
		return true
	}
}

func main() {
	var input string
	fmt.Scanln(&input)

	count := 0
	for i := 0; i < len(input); i++ {
		if isConsonant(input[i]) {
			count++
		}
	}

	fmt.Println(count)
}