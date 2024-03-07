package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if a == "A" {
		fmt.Print(5)
	} else if a == "B" {
		fmt.Print(4)
	} else if a == "C" {
		fmt.Print(3)
	} else if a == "D" {
		fmt.Print(2)
	} else if a == "E" {
		fmt.Print(1)
	} else if a == "T" {
		fmt.Print(0)
	}
}
