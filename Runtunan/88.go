package main

import "fmt"

func main() {
	var hari, minggu, bulan, tahun int
	fmt.Scan(&hari)
	tahun = hari / 360
	bulan = hari % 360 / 30
	minggu = hari % 360 % 30 / 7
	hari = hari % 360 % 30 % 7
	fmt.Println(tahun, bulan, minggu, hari)
}
