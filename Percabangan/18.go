package main

import "fmt"

func main() {
	var input string
	var hasil bool
	fmt.Scan(&input)
	hasil = input == "ya"
	if hasil {
		fmt.Scan(&input)
		hasil = input == "ya"
		if hasil {
			fmt.Println("tidak pergi ke minimarket")
		} else {
			fmt.Println("pergi ke minimarket")
		}
	} else {
		fmt.Println("tidak pergi ke minimarket")
	}
}
