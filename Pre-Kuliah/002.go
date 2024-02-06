// Program NIM
package main

import "fmt"

func main() {
	var tahun_masuk, tanggal_lahir, NIM, AA, BB, CC int
	fmt.Scan(&tahun_masuk, &tanggal_lahir)
	AA = (tahun_masuk % 100) * 10000
	BB = (tanggal_lahir % 100) * 100
	CC = tanggal_lahir / 1000000
	NIM = 130000000 + AA + BB + CC
	fmt.Println(NIM)
}
