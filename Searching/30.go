package main

import (
	"fmt"
)

const N = 1000

func hereThere(T [N]int, total, inikah int) int {
	// Implementasi binary search untuk menemukan posisi nilai 'inikah' atau nilai setelahnya
	low, high := 0, total-1
	
	// Jika 'inikah' lebih besar dari nilai terakhir dalam array, kembalikan -1
	if inikah > T[high] {
		return -1
	}

	for low <= high {
		mid := (low + high) / 2
		if T[mid] == inikah {
			return T[mid]
		} else if T[mid] < inikah {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	// Jika 'inikah' tidak ditemukan, 'low' akan berada di indeks nilai setelah 'inikah'
	return T[low]
}

func main() {
	var T [N]int
	var total, inikah int

	// Membaca jumlah total elemen dalam array
	fmt.Scan(&total)
	
	// Membaca elemen-elemen dalam array
	for i := 0; i < total; i++ {
		fmt.Scan(&T[i])
	}
	
	// Membaca nilai 'inikah'
	fmt.Scan(&inikah)

	// Memanggil fungsi hereThere dan mencetak hasilnya
	result := hereThere(T, total, inikah)
	fmt.Println(result)
}