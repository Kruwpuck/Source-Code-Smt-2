package main

import "fmt"

const N = 1000

func sudahTerurut(T [N]int, total int) bool {
	for i := 1; i < total; i++ {
		if T[i] < T[i-1] {
			return false // Jika ditemukan elemen yang lebih kecil dari elemen sebelumnya, maka tidak terurut
		}
	}
	return true // Jika tidak ditemukan, maka terurut
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

	// Memanggil fungsi sudahTerurut dan mencetak hasilnya
	result := sudahTerurut(T, total)
	fmt.Println(result)
}