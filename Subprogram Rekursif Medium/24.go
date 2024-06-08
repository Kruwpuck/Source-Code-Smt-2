func decimalToQuarternary(n int) int {
	if n == 0 {
		return 0
	} else {
		return (n % 4) + decimalToQuarternary(n/4)*10
	}
}