package main

import "fmt"

func main() {
	var hasil string
	fmt.Scan(&hasil)
	if hasil == "a" || hasil == "i" || hasil == "u" || hasil == "e" || hasil == "o" || hasil == "A" || hasil == "I" || hasil == "U" || hasil == "E" || hasil == "O" {
		fmt.Print("vokal")
	} else {
		fmt.Print("konsonan")
	}
}
