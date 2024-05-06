package main

import "fmt"

type HasilGame struct {
	TotalKemunculanGanjil int
}

func game() int {
	var d1, d2, d3 int
	var hasil HasilGame

	fmt.Scan(&d1, &d2, &d3)
	for (d1 + d2 + d3) != 10 {
		if d1%2 != 0 && d2%2 != 0 && d3%2 != 0 {
			hasil.TotalKemunculanGanjil++
		}
		fmt.Scan(&d1, &d2, &d3)
	}

	return hasil.TotalKemunculanGanjil
}

func main() {

	fmt.Println(game())
}
