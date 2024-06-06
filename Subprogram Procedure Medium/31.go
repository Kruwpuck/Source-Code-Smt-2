func pemenang(n int) {
	/*  I.S terdefinisi bilangan bulat n yang menyatakan banyaknya babak
		F.S karakter 'A', 'B', atau 'C' yang menyatakan pemenang game dan total angkanya */
		var a, b, c, totA, totB, totC, i int
		
		for i = 1; i <= n; i++{
			fmt.Scan(&a, &b, &c)
			totA += a
			totB += b
			totC += c
		}
		if totA > totB && totA > totC{
			fmt.Print("A", " ", totA)
		}else if totB > totA && totB > totC{
			fmt.Print("B", " ", totB)
		}else if totC > totA && totC > totB{
			fmt.Print("C", " ", totC)
	}
	}