package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Inisialisasi jumlah total angka ganjil untuk masing-masing asisten
	var totalA, totalB, totalC int

	// Baca hasil lemparan dadu untuk setiap babak
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)

		// Tambahkan angka ganjil ke total masing-masing asisten
		if a%2 != 0 {
			totalA += a
		}
		if b%2 != 0 {
			totalB += b
		}
		if c%2 != 0 {
			totalC += c
		}
	}

	// Tentukan pemenang berdasarkan total angka ganjil terbesar
	if totalA > totalB && totalA > totalC {
		fmt.Printf("A %d\n", totalA)
	} else if totalB > totalA && totalB > totalC {
		fmt.Printf("B %d\n", totalB)
	} else {
		fmt.Printf("C %d\n", totalC)
	}
}