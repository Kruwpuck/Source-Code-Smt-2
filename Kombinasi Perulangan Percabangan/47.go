package main

import (
	"fmt"
)

func main() {
	var n int
	var buah string
	var countApel, countJeruk, countMangga int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&buah)
		switch buah {
		case "apel":
			countApel++
		case "jeruk":
			countJeruk++
		case "mangga":
			countMangga++
		}
	}

	fmt.Print(countApel, " ")
	fmt.Print(countJeruk, " ")
	fmt.Print(countMangga, " ")
}