package main

import (
	"fmt"
)

func main() {
	var jumlahHari int
	fmt.Scan(&jumlahHari)

	var kemampuanA, kemampuanB float64

	for i := 1; i <= jumlahHari; i++ {
		var yin, yang int
		fmt.Scan(&yin, &yang)

		if yin < yang {
			kemampuanB += 0.15
		} else if yang < yin {
			kemampuanA += 0.25
		}

		if i%3 == 0 {
			kemampuanA = kemampuanA - (kemampuanA * 0.05)
		}
	}

	fmt.Print(kemampuanA, kemampuanB)
}
