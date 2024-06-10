package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	countVowel := 0
	countConsonant := 0

	i := 0
	for input[i] != '.' && i < len(input) {
		char := input[i]
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			countVowel++
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			countConsonant++
		}
		i++
	}

	fmt.Println(countVowel, countConsonant)
}