func printPolaSegitiga(baris int, kolom int, n int, angka int) {
	/* I.S Terdefinisi nilai bilangan bulat baris, kolom, dan n
	F.S menampilkan pola angka yang berbentuk segitiga */

	if kolom <= n {
		if kolom <= n-baris {
			fmt.Print(" ")
		} else {
			fmt.Print(angka % 10)
			angka++
		}
		printPolaSegitiga(baris, kolom+1, n, angka)
	} else if baris < n {
		fmt.Println()
		printPolaSegitiga(baris+1, 1, n, angka)
	}
}