func cetakdigit(n int) {
	/* I.S. Terdefinisi sebuah bilangan bulat n.
	F.S. Tercetak tiap digit bilangan n diselingi dengan karakter dash (strip). */

	if n < 10 {
		fmt.Print(n)
		return
	}

	// Rekurens: cetak digit pertama, lalu panggil rekursi untuk digit selanjutnya
	cetakdigit(n / 10)
	fmt.Print("-", n%10)
}