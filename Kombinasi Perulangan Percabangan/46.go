package main

import (
	"fmt"
)

func main() {
	var belajarCount, mainCount int
	var activity string

	for {
		fmt.Scanln(&activity)

		if activity == "tidur" {
			break
		} else if activity == "belajar" {
			belajarCount++
		} else if activity == "main" {
			mainCount++
		}
	}

	fmt.Print(belajarCount, " ")
	fmt.Print(mainCount, " ")
}