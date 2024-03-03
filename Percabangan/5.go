package main

import "fmt"

func main() {
	var kapasitas, penumpang float64
	fmt.Scan(&kapasitas, &penumpang)
	if penumpang/kapasitas >= 0.50 && penumpang/kapasitas <= 0.75 {
		fmt.Println("berangkat")
	} else {
		fmt.Println("tidak berangkat")
	}

}
