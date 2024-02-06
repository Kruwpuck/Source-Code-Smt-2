package main

import (
	"fmt"
	"strconv"
)

func main() {
	var tahunMasuk, tanggalLahir int

	// Membaca masukan tahun masuk dan tanggal lahir
	fmt.Print("Masukkan tahun masuk kuliah (format YYYY): ")
	fmt.Scan(&tahunMasuk)
	fmt.Print("Masukkan tanggal lahir (format DDMMYYYY): ")
	fmt.Scan(&tanggalLahir)

	// Mendapatkan dua digit terakhir dari tahun masuk
	aa := tahunMasuk % 100

	// Mendapatkan dua digit terakhir dari tahun lahir
	bb := (tanggalLahir % 10000) / 100

	// Mendapatkan dua digit tanggal lahir
	cc := tanggalLahir % 100

	// Menggabungkan nilai aa, bb, dan cc untuk membentuk NIM
	nim, _ := strconv.Atoi(fmt.Sprintf("1302%02d%02d%02d", aa, bb, cc))

	// Menampilkan hasil NIM
	fmt.Println("NIM Mahasiswa:", nim)
}
