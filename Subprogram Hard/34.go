func cetakterbalikdigit(n int) {
	/* I.S. Terdefinisi sebuah bilangan bulat n.
	F.S. Tercetak terbalik tiap digit bilangan n diselingi dengan karakter dash (strip). */

	if n < 10 {
		fmt.Print(n)
		return
	}

	// Cetak digit terbalik, kemudian panggil rekursi untuk digit selanjutnya
	fmt.Print(n%10, "-")
	cetakterbalikdigit(n / 10)
}
