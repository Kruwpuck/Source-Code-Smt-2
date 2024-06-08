func decimalToBiner(n int) int {
	if n == 0 {
		return 0
	} else {
		return (n % 2) + decimalToBiner(n/2)*10
	}
}