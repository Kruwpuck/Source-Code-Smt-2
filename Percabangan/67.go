package main

import "fmt"

func main() {
	var d1, d2, d3 int
	var hasil bool
	fmt.Scan(&d1, &d2, &d3)
	hasil = (d1-d2 == 1 && d2-d3 == 1) || (d3-d2 == 1 && d2-d1 == 1)
	fmt.Println(hasil)
}
