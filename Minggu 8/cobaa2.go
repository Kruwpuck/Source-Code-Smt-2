package main

import "fmt"

const NMAX int = 5

type matrix [NMAX][NMAX]int

func main() {
	var m1, m2, m3 matrix
	var n int

	fmt.Scan(&n) // ukuran matriks

	fmt.Println("Matriks pertama")
	baca(&m1, &n)
	cetak(m1, n)

	fmt.Println("Matriks kedua")
	baca(&m2, &n)
	cetak(m2, n)

	fmt.Println("Matriks ketiga")
	kurang(m1, m2, &m3, n)
	cetak(m3, n)
}

func baca(m *matrix, n *int) {
	/*
		IS: Matriks m terdefinisi sembarang, n terdefinisi
		FS: Matriks m dengan baris n x kolom n berisi nilai. Jika n > NMAX
		    gunakan nilai n = NMAX
	*/
	var i, j int
	if *n > NMAX {
		*n = NMAX
	}
	for i = 0; i < *n; i++ {
		for j = 0; j < *n; j++ {
			fmt.Scan(&m[i][j])
		}

	}
}

func cetak(m matrix, n int) {
	/*
		IS: Matrik m terdefinisi, n terdefinisi
		FS: Tercetak matriks m di layar dengan format:

			x11 x12 ... x1n
			x21 x22 ... x2n
			... ... ... ...
			xn1 xn2 ...	xnn

			dengan n adalah ukuran matriks
	*/
	var i, j int
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println()
	}
}

func kurang(A, B matrix, C *matrix, n int) {
	/*
		IS: Matriks A dan B terdefinisi
		FS: Matriks C berisi nilai. Elemen ke-i dan j matriks C merupakan
		    pengurangan elemen ke-i dan j dari matriks A oleh B.

			Contoh matriks dengan n = 2:
			matriks A - matriks B = matriks C
			1 1		 -   3 0	  =  -2 1
			2 2			 4 1		 -2 1
	*/
	var i, j int
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
}
