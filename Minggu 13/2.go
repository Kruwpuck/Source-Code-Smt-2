package main

import "fmt"

const nmax int = 37

type tHimpunan struct {
	anggota [nmax - 1]int
	panjang int
}

func main() {
	var a, b tHimpunan
	fmt.Print("Anggota himpunan 1: ")
	bacaMasukan(&a)
	fmt.Print("Anggota himpunan 2: ")
	bacaMasukan(&b)
	urut(&a)
	urut(&b)
	fmt.Println("Himpunan 1 = Himpunan 2?",sama(a, b))
}

func bacaMasukan(A *tHimpunan) {
	var i int
	var check bool
	check = true
	for i < nmax && check {

		fmt.Scan(&A.anggota[i])
		check = ada(*A, i)
		i += 1
	}
	A.panjang = i-1
}

func urut(A *tHimpunan) {
	var i, j int
	var key int
	for i = 1; i < A.panjang; i++ {
		key = A.anggota[i]
		j = i - 1
		for j >= 0 && A.anggota[j] > key {
			A.anggota[j+1] = A.anggota[j]
			j = j - 1
		}
		A.anggota[j+1] = key
	}

}
func sama(A, B tHimpunan) bool {
	var equal bool = true
	for i := 0; i < A.panjang; i++ {
		if A.anggota[i] != B.anggota[i] {
			equal = false
		}
	}
	return equal
}
func ada(A tHimpunan, i int)bool  {
	var check bool = true
	e := 0
	for e < i {
		if A.anggota[i] == A.anggota[e] {
			check = false
			e += nmax

		}
		e++
	}
	return check
}