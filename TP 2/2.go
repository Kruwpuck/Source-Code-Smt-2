package main

import "fmt"

func main() {
	var i, j, bil int
	fmt.Scan(&bil)

	for i = 1; i <= bil; i++ {
		for j = 1; j <= bil; j++ {
			if j == i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}
}
