package main

import "fmt"

func main() {
	var n, m, jumlah float64
	fmt.Scan(&n)
	for i := 0; i < int(n); i++ {
		fmt.Scan(&m)
		jumlah += m
	}
	fmt.Println(jumlah)
	fmt.Printf("%.2f", float64(jumlah/n))
}
