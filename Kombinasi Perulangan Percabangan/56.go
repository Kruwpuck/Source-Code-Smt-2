package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	// Buat map untuk mencatat jumlah masing-masing jenis suara
	suara := make(map[string]bool)

	// Membaca dan mencatat jenis suara yang dikeluarkan
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanln(&s)
		suara[s] = true
	}

	// Cek apakah burung dapat menirukan kelima jenis suara
	canImitate := suara["JA"] && suara["BE"] && suara["CI"] && suara["JE"] && suara["LO"]
	fmt.Println(canImitate)
}