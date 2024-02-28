package main

import "fmt"

func main() {
	var tahun int
	var hasil bool
	fmt.Scan(&tahun)
	if tahun%400 == 0 {
		hasil = true
	} else if tahun%100 == 0 {
		hasil = false
	} else if tahun%4 == 0 {
		hasil = true
	} else {
		hasil = false
	}
	fmt.Println(hasil)
}
