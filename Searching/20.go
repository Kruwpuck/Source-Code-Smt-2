
const N = 1000

func lebihBesar(T [N]int, jumlah, target int) int {
	count := 0
	for i := 0; i < jumlah; i++ {
		if T[i] > target {
			count++
		}
	}
	return count
}