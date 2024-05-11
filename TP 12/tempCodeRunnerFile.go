
func bacaData(A *tabInt, n *int) {
	/*
		IS: Array bilangan bulat A dan n terdefinisi sembarang
		FS: Array bilangan bulat A terisi data sebanyak n elemen.
			Elemen array terisi bilangan bulat menaik (ascending).
			Jika bilangan yang dibaca lebih kecil dari
			bilangan sebelumnya, data tidak masuk ke dalam array dan
			pembacanaan terhenti.
	*/
	var x int
	fmt.Scan(&x)
	A[0] = x
	*n = 1
	for i := 1; i < NMAX; i++ {
		fmt.Scan(&x)
		if x < A[i-1] {
			break
		}
		A[i] = x
		*n++
	}
}