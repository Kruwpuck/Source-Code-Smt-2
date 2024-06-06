func cetakHitungJumlah(N int, jum *int) {
	*jum = 0
	for i := 1; i <= N; i++ {
		fmt.Println(i)
		*jum += i
	}
}