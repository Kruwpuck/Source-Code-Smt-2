package main

import "fmt"

const nmax int = 37

type tHimpunan struct {
	anggota [nmax - 1]int
	panjang int
}

func main() {
	var a, b tHimpunan
	var i int
	fmt.Print("himpunan 1: ")
	baca(&a)
	fmt.Print("himpunan 2: ")
	baca(&b)

	selection_sort(&a)
	selection_sort(&b)
	for i = 0; i < a.panjang; i++ {
		fmt.Print(a.anggota[i], " ")
	}
	fmt.Println()
	for i := 0; i < b.panjang; i++ {
		fmt.Print(b.anggota[i], " ")
	}

	fmt.Println()
	fmt.Println(cek_sama(a, b))
}

func baca(A *tHimpunan) {
	var i, e int
	var check bool
	check = true
	for i < nmax && check {

		fmt.Scan(&A.anggota[i])
		e = 0
		for e < i {
			if A.anggota[i] == A.anggota[e] {
				check = false
				e += nmax

			}
			e++
		}
		i += 1
	}
	A.panjang = i-1
}

func selection_sort(A *tHimpunan) {
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
func cek_sama(A, B tHimpunan) bool {
	var sama bool = true
	for i := 0; i < A.panjang; i++ {
		if A.anggota[i] != B.anggota[i] {
			sama = false
		}
	}
	return sama
}