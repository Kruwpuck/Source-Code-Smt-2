type pemain struct {
	nama string
	poin int
}

const NMAX int = 1024

type dataPemain [NMAX]pemain

func isiArray(a *dataPemain, n int) {
	// algoritma untuk mengisi data pemain
	// hint : tambahkan variabel lokal dan input scan 
var i int
	i = 1
	for i <= n {
		fmt.Scan(&a[i].nama, &a[i].poin)
		i += 1
	}
}

func selectionSort(a *dataPemain, n int) {
	//algoritma selection sort secara ascending
	var pass, i, idx int
	var temp pemain
	pass = 2
	for pass <= n {
		i = pass
		idx = pass - 1
		for i <= n {
			if a[idx].poin > a[i].poin {
				idx = i
			}
			i += 1
		}
		temp = a[pass-1]
		a[pass-1] = a[idx]
		a[idx] = temp
		pass += 1
	}
}

func showArray(a dataPemain, n int) {
	// algoritma untuk menampilkan data pemain
	var i int
	i = 1
	for i <= n {
		fmt.Println(a[i].nama, a[i].poin)
		i += 1
	}
}