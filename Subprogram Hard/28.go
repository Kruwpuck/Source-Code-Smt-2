func secMax(max, secondMax *int) {
	/* I.S. terdefinisi bilangan terbesar max, bilangan terbesar kedua secondMax.
	   dan inputan barisan bilangan bulat yang berakhir jika bilangan bulat sama dengan 0
	   F.S. max adalah bilangan bulat terbesar dan secondMax adalah bilangan bulat terbesar kedua */

	var num, first, second int
	first = *max
	second = *secondMax

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > first {
			second = first
			first = num
		} else if num > second {
			second = num
		}
	}

	*max = first
	*secondMax = second
}
