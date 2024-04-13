func isiArray(n int) [100]int {
	/* mengembalikan array yang berisi bilangan bulat */
	var aray [100]int
	for i := 0; i < n; i++ {
		fmt.Scan(&aray[i])
	}
	return aray
}

func countPositiveSumNegative(n [100]int) (int, int) {
	/* mengembalikan jumlah bilangan positif dan negatif dari array */
	var pos, neg, i int
	for i = 0; i < len(n); i++ {
		if n[i] < 0 {
			neg += n[i]
		} else if n[i] > 0 {
			pos++
		}
	}
	return pos, neg
}
