func lowercase(s string, n int) {
	/*{ I.S. Terdefinisi s sebagai input string kapital
	F.S. menampilkan string yang semua huruf kecil menggunakan fungsi rekursif*/
	if n > 0 {
	lowercase(s, n-1) // masukkan fungsi rekursif pada baris ini
	fmt.Print(string(s[n-1] + 32))
	}
   }