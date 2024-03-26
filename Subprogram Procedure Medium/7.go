//buatlah procedure saja

func jumlahGanjil(N int, hasil *int) {
	for i := 1; i <= N; i++ {
		if i%2 != 0 {
			*hasil += i
		}
	}
}