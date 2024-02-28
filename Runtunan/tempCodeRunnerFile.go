package main

import "fmt"

func main() {
	var tahun int
	var hasil bool
	fmt.Scan(&tahun)
	hasil = (tahun%400 == 0 || tahun%4 == 0) && tahun%100 != 0
	fmt.Println(hasil)
}
