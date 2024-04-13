package main

import "fmt"

func main() {
	var input, hasil, hasil2 string
	fmt.Scan(&input)
	hasil = string(input[0])
	hasil2 = string(input[len(input)-1])
	if (hasil == "a" || hasil == "i" || hasil == "u" || hasil == "e" || hasil == "o") && (hasil2 == "a" || hasil2 == "i" || hasil2 == "u" || hasil2 == "e" || hasil2 == "o") {
		fmt.Print("vokal di awal dan di akhir")
	} else if hasil2 == "a" || hasil2 == "i" || hasil2 == "u" || hasil2 == "e" || hasil2 == "o" {
		fmt.Print("vokal di akhir")
	} else if hasil == "a" || hasil == "i" || hasil == "u" || hasil == "e" || hasil == "o" {
		fmt.Print("vokal di awal")
	} else {
		fmt.Print("tidak ada vokal")
	}
}
