func cetak(n, m int) {
	/*
	I.S. terdefinisi dua buah bilangan bulat positif n dan m, dengan n < m
	F.S. menampilkan barisan bilangan dari n hingga m
	*/
		var i int
		for i = n; i <= m; i++{
			fmt.Print(i, " ")
		}
		if n > m{
		   for i = m; i <= n; i++{
			   fmt.Print(i, " ")
		   } 
		}
	}