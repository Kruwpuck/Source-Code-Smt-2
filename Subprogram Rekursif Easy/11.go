func capitalize(s string, n int) {
	/*{ I.S. Terdefinisi s sebagai input string
	F.S. menampilkan string yang semua hurufnya kapital menggunakan fungsi rekursi */
	if n > 0 {
	capitalize(s, n-1) // masukan fungsi rekursif pada baris ini
	fmt.Print(string(s[n-1] - 32))
	}
   }