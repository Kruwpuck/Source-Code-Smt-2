package main

import "fmt"

const NMAX int = 100

type tabInt [NMAX]struct {
	nama     string
	nim      string
	eprt     int
	semester int
	ipk      float64
}

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	fmt.Println(tertinggi(data, nData), terendah(data, nData), rerata(data, nData))
}
func baca(A *tabInt, N *int) {
	i := 0
	for i < NMAX {
		fmt.Scan(&A[i].nama, &A[i].nim)
		if A[i].nim == "none" {
			break
		}
		fmt.Scan(&A[i].eprt, &A[i].semester, &A[i].ipk)
		i++
	}
	*N = i
}

func tertinggi(A tabInt, N int) int {
	var max = A[0].eprt
	for i := 1; i < N; i++ {
		if max < A[i].eprt {
			max = A[i].eprt
		}
	}
	return max
}
func terendah(A tabInt, N int) float64 {
	var min = A[0].ipk
	for i := 1; i < N; i++ {
		if min > A[i].ipk {
			min = A[i].ipk
		}
	}
	return min
}
func rerata(A tabInt, N int) float64 {
	var jumlah float64

	for i := 0; i < N; i++ {
		jumlah += float64(A[i].semester)
	}
	return jumlah / float64(N)
}
