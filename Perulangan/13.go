package main

import "fmt"

func main() {
	var n, m, jumlah float64
	fmt.Scan(&n, &m)
	for i := n; i <= m; i++ {
		jumlah += 4.0 / (i)
	}
	fmt.Printf("%.2f\n", jumlah)
}
