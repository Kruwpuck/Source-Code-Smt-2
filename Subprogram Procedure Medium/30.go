func skorDanPemenang(n int) {
	/*  I.S terdefinisi bilangan bulat n yang menyatakan banyaknya ronde
		F.S terdiri dari tiga nilai, yaitu skor petinju A, skor petinju B, dan karakter 'A' atau 'B' */
		var a, b, c, d, e, f, skorA, skorB, i int

		for i = 1; i <= n; i++{
		fmt.Scan(&a, &b, &c, &d, &e, &f)
		skorA = skorA + a + b + c
		skorB = skorB + d + e + f
		}
		if skorA > skorB{
			fmt.Print(skorA, skorB," ", "A")
		}else if skorB > skorA{
			fmt.Print(skorA, skorB, " ","B")
		}else if skorA == skorB{
			fmt.Print(skorA, skorB," ","0")
		}
}
	
        
    