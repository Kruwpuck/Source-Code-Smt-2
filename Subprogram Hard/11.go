func pembagianLoop(n, m int) int {
	quotient := 0
	for n >= m {
		n -= m
		quotient++
	}
	return quotient
}
