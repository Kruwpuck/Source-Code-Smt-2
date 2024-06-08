func jumlah(x, y int, hasil *int) {
	if x%2 != 0 && y%2 == 0 { // x ganjil dan y genap
		*hasil += x + y
	}
}