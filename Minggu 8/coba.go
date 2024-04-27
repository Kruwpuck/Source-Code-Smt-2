package main

import "fmt"

const NMAX int = 2021

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
}
func baca(A *tabInt, N *int) {
	i := 0
	for i < NMAX {
		fmt.Scan(&A[i])
		if A[i] < 0 {
			break
		}
		i++
	}
	*N = i
}

func tertinggi(A tabInt, N int) int {
	var max = A[0]
	for i := 1; i < N; i++ {
		if max < A[i] {
			max = A[i]
		}
	}
	return max
}

func cetak(A tabInt, N int) {
	fmt.Println(tertinggi(A, N))
}
