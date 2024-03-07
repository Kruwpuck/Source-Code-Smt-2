package main

import "fmt"

func main() {
	var masuk string
	var pecat int
	for i := 0; i < 5; i++ {
		fmt.Scan(&masuk)
		if masuk == "kalah" {
			pecat++
		}

	}
	if pecat == 5 {
		fmt.Println("dipecat")
	} else {
		fmt.Println("tidak dipecat")
	}
}
