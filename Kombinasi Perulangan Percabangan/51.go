package main

import (
	"fmt"
)

func main() {
	var jenis string
	var n, m int
	fmt.Scan(&jenis, &n, &m)

	if jenis == "ganjil" {
		for i := n; i <= m; i++ {
			if i%2 != 0 {
				fmt.Print(i, " ")
			}
		}
	} else if jenis == "genap" {
		for i := n; i <= m; i++ {
			if i%2 == 0 {
				fmt.Print(i, " ")
			}
		}
	} else {
		fmt.Println("Input tidak valid. Masukkan 'ganjil' atau 'genap'.")
	}
}