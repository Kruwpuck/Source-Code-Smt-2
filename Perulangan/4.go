package main

import "fmt"

func main() {
	var total, a, b float64
	fmt.Scan(&total, &a, &b)
	if a > b && a+b/total >= 0.80 {
		fmt.Println("Kandidat A menang")
	} else if a < b && (a+b)/total >= 0.80 {
		fmt.Println("Kandidat B menang")
	} else {
		fmt.Println("Tidak ada pemenang")
	}
}
