
const nMax int = 100

type dataNilai [nMax]int

func isiArray(a int, b *dataNilai) {
	if a%2 == 0 {
		for i := 0; i < a; i++ {
			fmt.Scan(&b[i])
		}
		cekKecurangan(a, *b)
	} else {
		fmt.Println("Jumlah array invalid !!!")
	}

}

func cekKecurangan(n int, m dataNilai) {
	var s1, s2 float64
	for i := 0; i < n; i++ {
		if i < n/2 {
			s1 += float64(m[i])
		} else {
			s2 += float64(m[i])
		}
	}

	rata1 := avgSesi1(s1, float64(n)/2.0)
	rata2 := avgSesi2(s2, float64(n)/2.0)

	fmt.Println("Rata rata sesi 1 :", rata1)
	fmt.Println("Rata rata sesi 2 :", rata2)

	if rata2-rata1 >= 20 {
		fmt.Println("Kemungkinan terjadi kecurangan")
	} else {
		fmt.Println("Tidak terjadi kecurangan")
	}
}

func avgSesi1(a float64, b float64) float64 {
	return float64(a) / float64(b)
}

func avgSesi2(a float64, b float64) float64 {
	return float64(a) / float64(b)
}
