package main

import "fmt"

func main() {
	var a string
	var b int
	for true {
		fmt.Scan(&a)
		if a == "." {
			break
		} else if a == "a" || a == "i" || a == "u" || a == "e" || a == "o" {
			b++
		}
	}
	fmt.Print(b)
}
