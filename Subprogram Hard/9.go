func cariMedian(a, b, c int) int {
	if (a > b && a < c) || (a < b && a > c) {
		return a
	} else if (b > a && b < c) || (b < a && b > c) {
		return b
	} else {
		return c
	}
}
