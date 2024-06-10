const N = 1000

func findMountainPeak(T [N]int, jumlah int) int {
	if jumlah < 3 {
		return -1
	}

	for i := 1; i < jumlah-1; i++ {
		if T[i] > T[i-1] && T[i] > T[i+1] {
			return i
		}
	}
	return -1
}