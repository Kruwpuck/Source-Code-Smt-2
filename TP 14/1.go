package main

import "fmt"

/*
Tipe bentukan struktur tDrakor dengan atribut:
	judul (string), rating (real), jumlah episode (integer),
	dan durasi (integer)
*/
type tDrakor struct {
	judul        string
	rating      float64
	jumlahEpisode, durasi int
}

// Konstanta NMAX dengan nilai 5
const NMAX int = 8

// Tipe alias tabDrakor untuk array of tDrakor dengan ukuran NMAX
type tabDrakor [NMAX]tDrakor

func main() {
	// Deklarasi variabel drakor bertipe array of tDrakor
	var A tabDrakor

	// Deklarasi variabel nDrakor bertipe integer untuk menampung banyaknya data drakor
	var nDrakor int

	// Pemanggilan prosedur bacaDataDrakor
	bacaDataDrakor(&A,&nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data sebelum diurutkan:")
	cetakDataDrakor(A,nDrakor)

	// Pemanggilan prosedur urutMenurun
	urutMenurun(&A,nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data setelah diurutkan menurun berdasarkan rating:")
	cetakDataDrakor(A, nDrakor)

	// Pemanggilan prosedur urutMenaik
	urutMenaik(&A,nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data setelah diurutkan menaik berdasarkan durasi:")
	cetakDataDrakor(A,nDrakor)
}

func bacaDataDrakor(A *tabDrakor, n *int) {
	/*
		IS: Array A sebanyak n terdefinisi sembarang
		Proses: 1) Membaca n. Nilai n tidak boleh melebihi NMAX.
				2) Membaca n data array A dengan atribut-atributnya.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	fmt.Scan(&*n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].judul,&A[i].rating,&A[i].jumlahEpisode,&A[i].durasi)
	}
}

func cetakDataDrakor(A tabDrakor, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen
		FS: Tercetak di layar data array A dengan format
			" Judul    Rating       Jum Ep   Durasi
			 <judul> <rating> <jum episode> <durasi>
			  ...	  ...		...			 ...."
			  Gunakan format string dengan kolom 20, 6, 6, dan 6
			  untuk masing-masing judul, rating, jumlah episode,
			  dan durasi. 
	*/
    fmt.Printf("%20s %6s %6s %6s\n", "Judul", "Rating", "Jum Ep", "Durasi")
    for i := 0; i < n; i++ {
        fmt.Printf("%20s %6.1f %6d %6d\n", A[i].judul, A[i].rating, A[i].jumlahEpisode, A[i].durasi)
    }
	fmt.Println()
}

func urutMenaik(A *tabDrakor, n int) {
	/*
		IS: Terdefinisi Array A sebanyak n elemen
		FS: Array A sebanyak n elemen terurut menaik berdasarkan
			durasi dengan menggunakan algoritma insertion sort
	*/
	var pass,i int
	var temp tDrakor
	pass = 1
	for pass <= n-1{
		i = pass
		temp = A[pass]
		for i>0 && temp.durasi < A[i-1].durasi {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func urutMenurun(A *tabDrakor, n int) {
	/*
		IS: Terdefinisi Array A sebanyak n elemen
		FS: Array A sebanyak n elemen terurut menurun berdasarkan
			rating dengan menggunakan algoritma selection sort
	*/
	var i,pass,idxMin int
	for pass = 0;pass <= n-2;pass++{
		idxMin=pass
		for i = pass+1;i<=n-1;i++{
			if A[i].rating > A[idxMin].rating{
				idxMin=i
			}
		}
		A[pass],A[idxMin]=A[idxMin],A[pass]
	}
}