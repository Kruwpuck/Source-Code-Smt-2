func pemenang(t1, t2, t3 int) {
	/*  I.S terdefinisi 3 angka pertama t1, t2, dan t3
		F.S menampilkan tiga nilai yang menyatakan berapa kali masing-masing mahasiswa itu menjadi pemenang */
		
		var tot1, tot2, tot3 int
		for t1 != 0 && t2 != 0 && t3 != 0{
			if t1 != t2 && t1 != t3 && t2 == t3{
				tot1 += 1
			}else if t3 != t2 && t3 != t1 && t2 == t1{
				tot3 += 1
			}else if t2 != t1 && t2 != t3 && t1 == t3{
				tot2+= 1
			}
			fmt.Scan(&t1, &t2, &t3)
		}
		fmt.Print(tot1, tot2, tot3)
	}

