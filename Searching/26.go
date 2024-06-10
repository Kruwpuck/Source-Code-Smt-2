
func nonKapital(kata string) int {
	count := 0
	for _, char := range kata {
		if char >= 'a' && char <= 'z' {
			count++
		}
	}
	return count
}