package main

import "fmt"

func main() {
	var x, digit1, digit2 int
	var berhenti, hasil1, hasil2 bool
	fmt.Scan(&x)
	if x == 0 {
		fmt.Print(true)
	} else {
		hasil1 = true
		berhenti = false
		for !berhenti {
			digit1 = x % 10
			x = x / 10
			digit2 = x % 10
			hasil2 = digit1-digit2 == 1 || digit2-digit1 == 1
			hasil1 = hasil1 && hasil2
			berhenti = x == digit2
		}
		fmt.Println(hasil1)
	}

}
