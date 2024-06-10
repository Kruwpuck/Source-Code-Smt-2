package main

import "fmt"

const NMax = 1000

type tabT [NMax]int

func main() {
	var n, i int
	var DM, M, x, y tabT

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&DM[i], &M[i], &x[i], &y[i])
	}

	for i = 0; i < n; i++ {
		ramuan := min(DM[i]/x[i], M[i]/y[i])
		fmt.Printf("total ramuan yang berhasil dibuat dari lemari persediaan %d sebanyak: %d\n", i+1, ramuan)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}