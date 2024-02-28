package main

import "fmt"

func main() {
	var d1, d2, d3 int
	var hasil bool
	fmt.Scan(&d1, &d2, &d3)
	hasil = d1%2 != 0 && d2%2 != 0 && d3%2 != 0
	fmt.Println(hasil)
}
