func lowerUpper(s string, n int) {
	/*{ I.S. Terdefinisi s sebagai input string dengan variasi case
	F.S. menampilkan string yang huruf kecilnya sudah diubah menjadi huruf kapita */
   
	if n > 0 { // gunakan operator yang tepat
	lowerUpper(s, n-1) // masukan fungsi rekursif pada barus ini
	if s[n-1] < 97 {
	fmt.Print(string(s[n-1] + 32)) // tampilkan string
	} else {
	fmt.Print(string(s[n-1] - 32)) // tampilkan string
	}
	}
   }
   