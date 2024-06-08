func decimalToOktal(n int) string {
	if n < 8 {
		return fmt.Sprint(n)
	}
	return decimalToOktal(n/8) + fmt.Sprint(n%8)
}