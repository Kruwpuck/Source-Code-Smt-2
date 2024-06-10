package main

import "fmt"

func main() {
	var pemenang [3]int

	for {
		var a, b, c int
		fmt.Scan(&a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}

		// Cek apakah ada pemenang pada putaran ini
		if (a != b && b == c) || (b != a && a == c) || (c != a && a == b) {
			if a != b && a != c {
				pemenang[0]++
			} else if b != a && b != c {
				pemenang[1]++
			} else {
				pemenang[2]++
			}
		}
	}

	fmt.Println(pemenang[0], pemenang[1], pemenang[2])
}