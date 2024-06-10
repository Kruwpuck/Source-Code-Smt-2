package main

import "fmt"

const N = 1000

func countThis(T [N]int, total, ini int) int {
	count := 0
	for i := 0; i < total; i++ {
		if T[i] == ini {
			count++
		} else if T[i] > ini {
			break // Karena array sudah terurut, berhenti loop jika nilai melebihi 'ini'
		}
	}
	return count
}

func main() {
	var T [N]int
	var total, ini int

	// Membaca jumlah total elemen dalam array
	fmt.Scan(&total)

	// Membaca elemen-elemen dalam array
	for i := 0; i < total; i++ {
		fmt.Scan(&T[i])
	}

	// Membaca nilai 'ini'
	fmt.Scan(&ini)

	// Memanggil fungsi countThis dan mencetak hasilnya
	result := countThis(T, total, ini)
	fmt.Println(result)
}