const N int = 1000

func searchBil(T [N]int, x int, n int) int {
	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right-left)/2

		if T[mid] == x {
			return mid
		}

		if T[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}