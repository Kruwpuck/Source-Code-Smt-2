package main

import "fmt"

const NMAX int = 6

type tabInt [NMAX]int

func main() {
	var data = tabInt{2, 7, 4, 1, 5, 3}
	var n int = 6
	bubble(&data, n)
	cetak(data, n)
}
func bubble(A *tabInt, n int) {
	var i, pass int
	for pass = 1; pass < n-1; pass++ {
		for i = 0; i <= n-1-pass; i++ {
			if A[i] > A[i+1] {
				swap(&A[i], &A[i+1])
			}
		}
	}
}
func swap(A, B *int) {
	var temp int
	temp = *A
	*A = *B
	*B = temp
}
func cetak(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}
