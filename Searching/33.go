package main

import "fmt"

const N = 1000

func yangTertukar(T [N]int, total int) int {
	count := 0

	for i := 1; i < total; i++ {
		if T[i] < T[i-1] {
			count++ // Jika nilai pada indeks ke-i lebih kecil dari nilai pada indeks ke-(i-1), tambahkan count
		}
	}

	return count
}

func main() {
	var T [N]int
	var total int

	// Membaca jumlah total elemen dalam array
	fmt.Scan(&total)

	// Membaca elemen-elemen dalam array
	for i := 0; i < total; i++ {
		fmt.Scan(&T[i])
	}

	// Memanggil fungsi yangTertukar dan mencetak hasilnya
	result := yangTertukar(T, total)
	fmt.Println(result)
}