const nMax int = 1000
type tabInt [nMax]int
func isiArray(arrInt *tabInt, n *int) {
	var bil int
	fmt.Scan(&bil)
	for bil >= 0 && *n < nMax {
		(*arrInt)[*n] = bil
		*n++
		fmt.Scan(&bil)
	}
}

func checkTripletSum(arrInt tabInt, n, x int) {
	/* I.S. array arrInt berisi sejumlah n bilangan bulat positif
	   F.S. 3 bilangan yang jumlahnya sama dengan bilangan X telah
	   ditampilkan di layar */
	var found bool
	var i, j, k, sum int
	for i = 0; i < n-2; i++ {
		for j = i + 1; j < n-1; j++ {
			for k = j + 1; k < n; k++ {
				sum = arrInt[i] + arrInt[j] + arrInt[k]
				if sum == x {
					fmt.Println(arrInt[i], arrInt[j], arrInt[k])
					found = true
					// Memberi tanda bahwa triplet sudah ditemukan
					// untuk menghindari penampilan ganda
					arrInt[i] = -1
					arrInt[j] = -1
					arrInt[k] = -1
				}
			}
		}
	}

	if !found {
		fmt.Println("-")
	}
}