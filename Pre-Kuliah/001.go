// Program Konversi Waktu ke Detik
package main

import "fmt"

func main() {
	var jam, menit, detik, hasil int
	fmt.Scan(&jam, &menit, &detik)
	hasil = (jam * 3600) + (menit * 60) + detik
	fmt.Println(hasil)
}
