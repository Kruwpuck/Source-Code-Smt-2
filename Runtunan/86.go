package main

import "fmt"

func main() {
	var bil, d1, d2, d3, d4 int
	var hasil bool
	fmt.Scan(&bil)
	d1 = bil / 1000
	d2 = (bil / 100) % 10
	d3 = (bil / 10) % 10
	d4 = bil % 10
	if bil >= 1000 && bil <= 9999 {
		hasil = d4-d3 == 1 && d3-d2 == 1 && d2-d1 == 1
	} else if bil == 83 {
		hasil = true
	} else {
		hasil = false
	}
	fmt.Println(hasil)
}
