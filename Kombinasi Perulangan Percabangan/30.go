package main

import "fmt"

func main() {
	var input, hasil string
	for i := 0; i < 6; i++ {
		fmt.Scan(&input)
		hasil += string(input[0])
	}
	if hasil == "serang" {
		fmt.Println(hasil)
	} else {
		fmt.Println("tidak serang")
	}
}
