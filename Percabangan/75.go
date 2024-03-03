package main

import "fmt"

func main() {
	var bil, d1, d4 int
	var hasil bool
	fmt.Scan(&bil)
	d1 = bil / 1000
	d4 = bil % 10
	hasil = d1 == d4
	fmt.Println(hasil)
}
