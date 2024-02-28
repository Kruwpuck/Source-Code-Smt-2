package main

import "fmt"

func main() {
	var harga int
	var cashb bool
	fmt.Scan(&harga)
	cashb = harga >= 0 && harga <= 100
	fmt.Println(cashb)
}
