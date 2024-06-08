func modulusLoop(n, m int) int {
	for n >= m {
		n -= m
	}
	return n
}