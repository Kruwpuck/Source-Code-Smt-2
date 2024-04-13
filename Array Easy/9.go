const NMax int = 100

type tabInt [NMax]int

func masukanArray(T *tabInt, n *int) {
	/*{I.S. Data tersedia dalam piranti masukan
	  F.S. Array T berisi sejumlah n bilangan bulat yang dimasukka user.
	  Proses input akan berhenti ketika pengguna memasukkan angka -1717}
	  Petunjuk: Gunakan while loop untuk proses input*/
	var i int
	for i = 0; i < NMax; i++ {
		fmt.Scan(&T[i])
		if T[i] == -1717 {
			break
		}
	}
	*n = i
}

func filterArray(T1 tabInt, n1 int, T2 *tabInt, n2 *int, ambangBatas int) {
	/*{I.S. Terdefinisi array T1 yang berisikan n1 jumlah bilangan bulat
	  F.S. Array T2 berisikan elemen T1 yang <= ambangBatas)}*/
	var j, x int
	x = 0
	for i := 0; i < n1; i++ {
		if T1[i] <= ambangBatas {
			T2[x] = T1[i]
			j++
			x++
		}
	}
	*n2 = j
}

func cetakArray(T tabInt, n int) {
	/*{I.S. Terdefinisi array T berisikan n buah bilangan bulat
	  F.S. Seluruh elemen array T ditampilkan pada layar}*/
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
}