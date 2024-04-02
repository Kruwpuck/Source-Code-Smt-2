package main

import "fmt"

const NMAX int = 100

type TabTB [NMAX]int

func main() {
	var tb TabTB
	var n int
	var rerata float64

	fmt.Scan(&n)
	baca_data(&tb, n)
	rerata = rata_rata(tb, n)
	fmt.Println(rerata)
}
func baca_data(T *TabTB, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}
func rata_rata(T TabTB, n int) float64 {
	var rata float64
	var jumlah int
	jumlah = 0
	for i := 0; i < n; i++ {
		jumlah += T[i]
	}
	rata = float64(jumlah) / float64(n)
	return rata
}
func tertinggi(T TabMhs, n int) string {
	var tertinggi string
	var max int
	max = T[0].tb
	for i := 1; i < n; i++ {
		if max < T[i].tb {
			tertinggi = T[i].nama
		}
	}
	return tertinggi
}
