package main

import "fmt"

const NMAX = 2022

type mahasiswa struct {
	nama  string
	jenis byte
	IPK   float64
}

type tabMhs [NMAX]mahasiswa

func main() {
	var k tabMhs
	var n int
	fmt.Scan(&n)
	bacadata(&k, n)
	fmt.Println(femalestudent(k, n))
}

func bacadata(k *tabMhs, n int) {
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %c %f", &k[i].nama, &k[i].jenis, &k[i].IPK)
	}
}

func femalestudent(A tabMhs, n int) string {
	var idx int
	for i := 0; i < n; i++ {
		if A[i].jenis == 'f' && (idx == 0 || A[i].IPK > A[idx].IPK) {
			idx = i
		}
	}
	return A[idx].nama
}
