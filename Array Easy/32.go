const NMax int = 100
type tabUsia [NMax]int

func inputArray(T *tabUsia, n *int) {
	/*I.S. Data usia pasien tersedia dalam piranti masukan
	  Proses: Input akan terus berlangsung sampai bilangan yang diinput tidak
	  valid (<=0)
	  F.S. Sejumlah n data usia pasien masuk ke dalam array T
	  Petunjuk : Gunakan while loop untuk melakukan perulangan
	*/
	*n = 0
	for {
		var usia int
		fmt.Scan(&usia)
		if usia <= 0 {
			break
		}
		(*T)[*n] = usia
		*n++
	}
}

func hitungRata(T tabUsia, n int) float64 {
	/Mengembalikan rata-rata usia pasien pada array T/
	total := 0
	for i := 0; i < n; i++ {
		total += T[i]
	}
	return float64(total) / float64(n)
}