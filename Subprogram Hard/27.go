package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	bilanganPola(n, m)
}

func bilanganPola(n, m int) {
	/*
	IS : Terdefinisi bilangan bulat n dan m
	FS : menampilkan pola barisan bilangan
	*/
	for i := n; i <= m; i++ {
		if i%2 == 1 {
			// Jika sukunya ganjil, cetak 2^i
			fmt.Printf("%.2f ", math.Pow(2, float64(i)))
		} else {
			// Jika sukunya genap, cetak 2^-i
			if i < 0 {
				fmt.Printf("%.2f ", math.Pow(2, float64(i)))
			}else {
				fmt.Printf("%.2f ", math.Pow(2, float64(-i)))
			}
		}
	}
	fmt.Println()
}
