func pola(n int, s string) {
	/* I.S. Terdefinisi nilai bilangan bulat n dimana n>0 */
	/* F.S. prosedur akan menampilkan pola dengan string (*) */
	if n <= 0 {
		return
	}

	fmt.Println(s)
	pola(n-1, s+"*")
	if n > 1 {
		fmt.Println(s)
	}
}