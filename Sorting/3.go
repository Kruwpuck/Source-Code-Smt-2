type barang struct {
	nama  string
	poin int
}

const NMAX int = 1024

type dataBarang [NMAX]barang

func isiArray(a *dataBarang, n int) {
	// algoritma untuk menginput data barang dalam bentuk array 
	// hint : gunakan variabel lokal dan input scan
	var i int
	i = 0
	for i < n {
		fmt.Scan(&a[i].nama, &a[i].poin)
		i += 1
	}
}

func insertionSort(a *dataBarang, n int) {
	//algoritma insertion sort secara descending
	var pass, i int
	var temp barang
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = a[pass]
		for i > 0 && temp.poin > a[i-1].poin {
			a[i] = a[i-1]
			i -= 1
		}
		a[i] = temp
		pass++
	}
}

func showArray(a dataBarang, n int) {
	//algoritma menampilkan data barang
	var i int
	i = 0
	for i < n {
		fmt.Println(a[i].nama, a[i].poin)
		i += 1
	}
}