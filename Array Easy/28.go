const NMax int = 10000
type tabInt [NMax]int

func InputArray(Tab *tabInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&(*Tab)[i])
	}
}

func jumlahArrGenap(Tab tabInt, n int) int {
	
	var i, total int
	for i = 0; i < n; i++ {
		if Tab[i]%2 == 0 {
			total += Tab[i]
		}
	}
	return total
}