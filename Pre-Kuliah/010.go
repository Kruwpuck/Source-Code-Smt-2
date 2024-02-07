package main

import "fmt"

func main() {
	var x, bil, d1, d2, d3, d4, jumlah int
	stop := false
	fmt.Scan(&x)
	for !stop {
		fmt.Scan(&bil)
		d1 = bil / 1000
		d2 = bil / 100 % 10
		d3 = bil / 10 % 10
		d4 = bil % 10
		jumlah += d2 + d3
		stop = d1+d4 == x
	}
	fmt.Println(jumlah)
}
