package main

import "fmt"

func main() {
	var jam, menit, detik int
	fmt.Scan(&jam, &menit, &detik)
	fmt.Println(waktu(jam, menit, detik))
}
func waktu(jam, menit, detik int) int {
	var hasil int
	hasil = jam*3600 + menit*60 + detik
	return hasil
}
