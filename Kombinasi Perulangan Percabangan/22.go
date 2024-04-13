package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	hasil := string(input[0]) + string(input[4]) + string(input[8]) + string(input[12])

	if len(input) >= 20 {
		fmt.Println(hasil)
	}
}
