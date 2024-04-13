package main

import "fmt"

func main() {
	var suwir, cakue, ati, porsi bool
	var harga int
	fmt.Scan(&suwir, &cakue, &ati, &porsi)
	harga = 6000
	if suwir {
		harga += 3000
	}
	if cakue {
		harga += 1500
	}
	if ati {
		harga += 4500
	}
	if porsi {
		harga += 4000
	}
	fmt.Println(harga)
}
