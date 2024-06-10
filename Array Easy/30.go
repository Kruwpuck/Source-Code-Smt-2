package main

import "fmt"

const NMax int = 1000

type tabArr [NMax]int

func main() {
	var arr1, arr2 tabArr
	var n, i, result int

	// Input banyaknya array (n)
	fmt.Scan(&n)

	// Input elemen array pertama sebanyak n kali
	for i = 0; i < n; i++ {
		fmt.Scan(&arr1[i])
	}

	// Input elemen array kedua sebanyak n kali
	for i = 0; i < n; i++ {
		fmt.Scan(&arr2[i])
	}

	// Perulangan yang didalamnya terdapat perbandingan
	for i = 0; i < n; i++ {
		// Jika elemen array pertama dan kedua berbeda, lakukan operasi perkalian
		if arr1[i] != arr2[i] {
			result = arr1[i] * arr2[i]
			fmt.Print(result, " ")
		} else {
			// Jika sama, tampilkan elemen array pertama
			fmt.Print(arr1[i], " ")
		}
	}
}