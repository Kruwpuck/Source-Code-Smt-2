package main

import "fmt"

const NMax int = 1000

type tabG [NMax]int

func main() {
	var N, i, hasil int
	var K string
	var KA, KB tabG

	// Input banyaknya siswa
	fmt.Scan(&N)

	// Input nilai kelompok A sebanyak N
	for i = 0; i < N; i++ {
		fmt.Scan(&KA[i])
	}

	// Input nilai kelompok B sebanyak N
	for i = 0; i < N; i++ {
		fmt.Scan(&KB[i])
	}

	// Input kriteria yang ingin digunakan (ganjil/genap)
	fmt.Scan(&K)

	// Inisialisasi hasil penjumlahan
	hasil = 0

	// Menjumlahkan elemen dari 2 array dengan kriteria masukan
	if K == "genap" {
		// Jumlahkan elemen yang genap dari kedua array
		for i = 0; i < N; i++ {
			if KA[i]%2 == 0 {
				hasil += KA[i]
			}
			if KB[i]%2 == 0 {
				hasil += KB[i]
			}
		}
	} else if K == "ganjil" {
		// Jumlahkan elemen yang ganjil dari kedua array
		for i = 0; i < N; i++ {
			if KA[i]%2 != 0 {
				hasil += KA[i]
			}
			if KB[i]%2 != 0 {
				hasil += KB[i]
			}
		}
	}

	// Menampilkan hasil penjumlahan
	fmt.Printf("Hasil dari penjumlahan elemen array yang bernilai %s adalah %d\n", K, hasil)
}