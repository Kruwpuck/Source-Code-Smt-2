func cetakHitungJumlahRata(n, m int, jum *int, rata *float64) {
	*jum = 0
	*rata = 0.
	for i := n; i <= m; i++ {
	*jum += i
	fmt.Println(i)
	}
	*rata = float64(*jum) / float64(m-n+1)
   }