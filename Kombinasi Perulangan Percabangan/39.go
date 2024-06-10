package main

import (
	"fmt"
)

func main() {
	var x, y, z int

	fmt.Scan(&x, &y, &z)

	totalPertemuan := hitungPertemuan(x, y, z)
	fmt.Println(totalPertemuan)
}

func hitungPertemuan(x, y, z int) int {
	jumlahHari := 365
	pertemuan := 0

	for hari := 1; hari <= jumlahHari; hari++ {
		if hari%x == 0 && hari%y == 0 && hari%z != 0 {
			pertemuan++
		}
	}

	return pertemuan
}