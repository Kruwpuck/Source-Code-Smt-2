package main

import "fmt"

// Lengkapi fungsi penjumlahan di bawah
func penjumlahan(b1, b2 int) int {
// Mengembalikan nilai berupa penjumlahan b1 dan b2
	var hasil int
	hasil = b1 + b2
	return hasil
}
func main() {
	var bil1, bil2, hasil_penjumlahan int
	fmt.Scan(&bil1, &bil2) // input dalam program utama
	hasil_penjumlahan = penjumlahan(bil1, bil2) // Fungsi dipanggil dalam program u
	fmt.Println(hasil_penjumlahan) // output dalam program utama
}