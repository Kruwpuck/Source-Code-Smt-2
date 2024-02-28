package main

import (
	"fmt"
	"unicode"
)

func main() {
	var char1, char2 rune
	fmt.Scanf("%c %c", &char1, &char2)

	lowerChar1 := unicode.ToLower(char1)
	lowerChar2 := unicode.ToLower(char2)

	hasil := lowerChar1 == lowerChar2

	fmt.Println(hasil)
}
